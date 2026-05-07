import { NavLink } from "react-router-dom";
import styles from "./Navbar.module.css";
import uchiLogo from "../assets/logo/uchi.svg";

const Navbar = () => {
  return (
    <nav className={styles.navbar}>
      <div className={styles.inner}>
        <div className={styles.logo}>
          <img src={uchiLogo} alt="Uchi" className={styles.logoImage} />
          <span>ТАМАГОЧИ</span>
        </div>

        <div className={styles.links}>
          <NavLink
            to="/game"
            className={({ isActive }) =>
              isActive ? `${styles.link} ${styles.active}` : styles.link
            }
          >
            ИГРА
          </NavLink>

          <NavLink
            to="/books"
            className={({ isActive }) =>
              isActive ? `${styles.link} ${styles.active}` : styles.link
            }
          >
            КНИГИ
          </NavLink>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
