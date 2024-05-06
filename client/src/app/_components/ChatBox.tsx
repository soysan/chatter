import { FC, useEffect, useRef } from "react";
import styles from "./ChatBox.module.css"

export type Message = {
  user_input: string,
  bot_response?: string,
  response_timestamp?: string,
  request_timestamp?: string
}

type ChatBoxProps = {
  messages: Message[]
}

const ChatBox: FC<ChatBoxProps> = ({messages}) => {
  const scrollRef = useRef<HTMLDivElement>(null)
  useEffect(() => {
    scrollRef.current?.scrollIntoView({behavior: 'smooth'})
  }, [messages])

  return (
    <div className={styles.container}>
      {messages.map((msg, i) => (
        msg.request_timestamp ? (
          <div key={i} className={styles.userBox}>
            <p className={`${styles.msg} ${styles.userMsg}`}>{msg.user_input}</p>
            <span className={styles.time}>{msg.request_timestamp}</span>
          </div>
        ) : (
          <div key={i}>
            <p className={`${styles.msg} ${styles.botMsg}`}>{msg.bot_response}</p>
            <span className={styles.time}>{msg.response_timestamp}</span>
          </div>
        )
      ))
      }
      <div ref={scrollRef}/>
    </div>
  );
}

export default ChatBox
