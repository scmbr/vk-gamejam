// src/data/books.ts
export interface Book {
  id: string;
  title: string;
  author: string;
  difficulty: string;
  duration: string;
  description: string;
  pdfFile: string;
}

export const books: Book[] = [
  {
    id: "book-1",
    title: "Книга",
    author: "Тамагочи",
    difficulty: "Лёгкий",
    duration: "15 мин",
    description: "Сборника рассказов",
    pdfFile: "book-1.pdf",
  },
];
