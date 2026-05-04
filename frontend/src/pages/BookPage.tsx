import { useParams, useNavigate } from "react-router-dom";
import { books } from "../data/books";
import styles from "./BookPage.module.css";

const BookPage = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();

  const book = books.find((b) => b.id === id);

  if (!book) {
    return (
      <div className={styles.notFound}>
        <p>Книга не найдена</p>
        <button onClick={() => navigate("/books")}>← Назад</button>
      </div>
    );
  }

  const handleDownload = async () => {
    const response = await fetch(`/books/${book.pdfFile}`);
    const blob = await response.blob();
    const url = URL.createObjectURL(blob);
    const link = document.createElement("a");
    link.href = url;
    link.download = `${book.title}.pdf`;
    link.click();
    URL.revokeObjectURL(url);
  };

  return (
    <div className={styles.container}>
      <button className={styles.back} onClick={() => navigate("/books")}>
        ← Назад
      </button>
      <div className={styles.card}>
        <h1 className={styles.title}>{book.title}</h1>
        <button className={styles.downloadButton} onClick={handleDownload}>
          📥 Скачать PDF
        </button>
      </div>
    </div>
  );
};

export default BookPage;
