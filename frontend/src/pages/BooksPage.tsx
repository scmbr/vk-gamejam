import { useNavigate } from "react-router-dom";
import { books } from "../data/books";
import styles from "./BooksPage.module.css";

const BooksPage = () => {
  const navigate = useNavigate();

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Библиотека</h1>
      <div className={styles.grid}>
        {books.map((book) => (
          <div
            key={book.id}
            className={styles.card}
            onClick={() => navigate(`/books/${book.id}`)}
          >
            <div className={styles.icon}>📖</div>
            <div className={styles.info}>
              <h2 className={styles.bookTitle}>{book.title}</h2>
              <p className={styles.author}>{book.author}</p>
              <div className={styles.tags}>
                <span className={styles.tag}>{book.difficulty}</span>
                <span className={styles.tag}>{book.duration}</span>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default BooksPage;
