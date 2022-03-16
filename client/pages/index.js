import Head from "next/head";
import styles from "../styles/Home.module.css";
import React, { useEffect, useState } from "react";

const Home = () => {
  const [response, setResponse] = useState(String);

  useEffect(async () => {
    const hostname = window.origin;
    const res = await fetch(`${hostname}/api/ping`);
    const data = await res.json();
    console.log({ data });
    setResponse(data);
  }, []);

  return (
    <div className={styles.container}>
      <Head>
        <title>Create Next App</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}>
          Welcome to Go embed <a href="https://nextjs.org">Next.js!</a>
        </h1>
        <h2>Pong: {response.pong}</h2>
      </main>

      <footer className={styles.footer}>
        <a href="https://go.dev/" target="_blank" rel="noopener noreferrer">
          Powered by Go
        </a>
      </footer>
    </div>
  );
};

export default Home;
