import { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";
import styles from "./LoginPage.module.css";
import animals from "../assets/group.svg";

const RegisterPage = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirm, setConfirm] = useState("");
  const [localError, setLocalError] = useState("");
  const { register, loading, error } = useAuth();
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (password !== confirm) {
      setLocalError("Пароли не совпадают");
      return;
    }
    const token = await register(email, password);
    if (token) navigate("/game");
  };

  return (
    <div className={styles.container}>
      <div className={styles.card}>
        <div className={styles.logo}>
          <img src={animals} alt="Animals" className={styles.animals} />
        </div>
        <h1 className={styles.title}>Регистрация</h1>
        <form onSubmit={handleSubmit} className={styles.form}>
          <input
            className={styles.input}
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
          <input
            className={styles.input}
            type="password"
            placeholder="Пароль"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
          <input
            className={styles.input}
            type="password"
            placeholder="Повтори пароль"
            value={confirm}
            onChange={(e) => setConfirm(e.target.value)}
            required
          />
          {(error || localError) && (
            <p className={styles.error}>{localError || error}</p>
          )}
          <button className={styles.button} type="submit" disabled={loading}>
            {loading ? "Загрузка..." : "ЗАРЕГИСТРИРОВАТЬСЯ"}
          </button>
        </form>
        <p className={styles.link}>
          Уже есть аккаунт? <Link to="/login">Войти</Link>
        </p>
      </div>
    </div>
  );
};

export default RegisterPage;
