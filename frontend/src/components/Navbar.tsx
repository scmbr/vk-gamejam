import { NavLink } from "react-router-dom";
import styles from "./Navbar.module.css";

const Navbar = () => {
  return (
    <nav className={styles.navbar}>
      <div className={styles.logo}>Тамагочи</div>
      <div className={styles.links}>
        <NavLink
          to="/game"
          className={({ isActive }) =>
            isActive ? `${styles.link} ${styles.active}` : styles.link
          }
        >
          Игра
        </NavLink>
        <NavLink
          to="/books"
          className={({ isActive }) =>
            isActive ? `${styles.link} ${styles.active}` : styles.link
          }
        >
          Книги
        </NavLink>
      </div>
    </nav>
  );
};

export default Navbar;
