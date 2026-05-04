import { useEffect, useState } from "react";
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import { setAccessToken, getAccessToken } from "./api/client";
import client from "./api/client";
import type { AuthResponse } from "./types";
import Navbar from "./components/Navbar";
import LoginPage from "./pages/LoginPage";
import RegisterPage from "./pages/RegisterPage";
import GamePage from "./pages/GamePage";
import BooksPage from "./pages/BooksPage";
import BookPage from "./pages/BookPage";

const App = () => {
  const [isReady, setIsReady] = useState(false);

  useEffect(() => {
    const tryRestoreSession = async () => {
      try {
        const { data } = await client.post<AuthResponse>("/auth/refresh");
        setAccessToken(data.accessToken);
      } catch {
        setAccessToken("");
      } finally {
        setIsReady(true);
      }
    };
    tryRestoreSession();
  }, []);

  if (!isReady) return <div>Загрузка...</div>;

  const isAuth = !!getAccessToken();

  return (
    <BrowserRouter>
      {isAuth && <Navbar />}
      <Routes>
        <Route path="/login" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route
          path="/game"
          element={isAuth ? <GamePage /> : <Navigate to="/login" />}
        />
        <Route
          path="/books"
          element={isAuth ? <BooksPage /> : <Navigate to="/login" />}
        />
        <Route
          path="/books/:id"
          element={isAuth ? <BookPage /> : <Navigate to="/login" />}
        />
        <Route
          path="*"
          element={<Navigate to={isAuth ? "/game" : "/login"} />}
        />
      </Routes>
    </BrowserRouter>
  );
};

export default App;
