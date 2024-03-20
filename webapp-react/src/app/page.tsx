import Image from "next/image";
import styles from "./page.module.css";
import { Typography } from "@mui/material";

export default function Home() {
  return (
    <main className={styles.main}>
      <div className={styles.description}>
        <Typography variant="h1">Welcome to Kadai</Typography>
      </div>
    </main>
  );
}
