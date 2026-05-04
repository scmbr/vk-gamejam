import { useEffect, useRef } from "react";
import { useNavigate } from "react-router-dom";
import { getAccessToken, setAccessToken } from "../api/client";
import client from "../api/client";
import type { AuthResponse } from "../types";
import styles from "./GamePage.module.css";

declare global {
  interface Window {
    createUnityInstance: (
      canvas: HTMLCanvasElement,
      config: object,
      onProgress?: (progress: number) => void,
    ) => Promise<any>;
    unityInstance: any;
  }
}

const GamePage = () => {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const unityInstanceRef = useRef<any>(null);
  const navigate = useNavigate();

  useEffect(() => {
    const token = getAccessToken();
    if (!token) {
      navigate("/login");
      return;
    }

    // Unity сигналит что готова — отправляем токен
    const handleUnityReady = () => {
      const currentToken = getAccessToken();
      if (currentToken) {
        unityInstanceRef.current?.SendMessage(
          "TokenBridge",
          "SetAccessToken",
          currentToken,
        );
      }
    };

    // Unity просит refresh
    const handleRefreshRequest = async () => {
      try {
        const { data } = await client.post<AuthResponse>("/auth/refresh");
        setAccessToken(data.accessToken);
        unityInstanceRef.current?.SendMessage(
          "TokenBridge",
          "OnTokenRefreshed",
          data.accessToken,
        );
      } catch (error: any) {
        // Только если 401 — токен реально протух
        if (error.response?.status === 401) {
          navigate("/login");
        }
        // Иначе просто логируем — возможно race condition
        console.error("Refresh failed:", error);
      }
    };

    // beforeunload — надёжное сохранение при закрытии вкладки
    const handleBeforeUnload = () => {
      const currentToken = getAccessToken();
      if (!currentToken) return;

      // sendBeacon гарантированно отправляет даже при закрытии
      navigator.sendBeacon(
        "http://localhost:8000/api/session/end",
        new Blob([JSON.stringify({})], { type: "application/json" }),
      );
    };

    window.addEventListener("unity_ready", handleUnityReady);
    window.addEventListener("unity_request_refresh", handleRefreshRequest);
    window.addEventListener("beforeunload", handleBeforeUnload);

    const script = document.createElement("script");
    script.src = "/Build/Build.loader.js";
    script.onload = () => {
      if (canvasRef.current) {
        window
          .createUnityInstance(
            canvasRef.current,
            {
              dataUrl: "/Build/Build.data.br",
              frameworkUrl: "/Build/Build.framework.js.br",
              codeUrl: "/Build/Build.wasm.br",
              streamingAssetsUrl: "StreamingAssets",
              loaderUrl: "/Build/Build.loader.js",
            },
            (progress) => {
              console.log(`Loading: ${Math.round(progress * 100)}%`);
            },
          )
          .then((instance: any) => {
            unityInstanceRef.current = instance;
            window.unityInstance = instance;
          });
      }
    };
    document.body.appendChild(script);

    return () => {
      window.removeEventListener("unity_ready", handleUnityReady);
      window.removeEventListener("unity_request_refresh", handleRefreshRequest);
      window.removeEventListener("beforeunload", handleBeforeUnload);
      document.body.removeChild(script);
    };
  }, [navigate]);

  return (
    <div className={styles.container}>
      <canvas ref={canvasRef} className={styles.canvas} id="unity-canvas" />
    </div>
  );
};

export default GamePage;
