'use client'

import styles from "./page.module.css";
import InputForm from "@/app/_components/InputForm";
import ChatBox, { Message } from "@/app/_components/ChatBox";
import { useTime } from "@/app/_hooks/useTime";
import { useState } from "react";

export default function Home() {
  const [messages, setMessages] = useState<Message[]>([]);
  const {displayTime} = useTime()

  const handleMessages = async (newMsg: string) => {
    const url = "" + process.env.NEXT_PUBLIC_SERVER_URL
    const res = await fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({user_input: newMsg}),
    });

    if (!res.ok) {
      console.error('Failed to send message');
      return;
    }

    const userInput = {user_input: newMsg, request_timestamp: displayTime()}

    const data = await res.json();
    data.response_timestamp = displayTime(data.response_timestamp);
    setMessages([...messages, userInput, data]);
  }

  return (
    <main className={styles.main}>
      <h1 className={styles.title}>Chat bots!!!!!</h1>
      <ChatBox messages={messages}/>
      <InputForm handleMessages={handleMessages}/>
    </main>
  );
}
