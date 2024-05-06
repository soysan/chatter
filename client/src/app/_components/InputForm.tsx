'use client'

import { FC, FormEvent, FormEventHandler, useRef } from 'react'
import styles from './InputForm.module.css'

type InputFormProps = {
  handleMessages: (newMsg: string) => void
}

const InputForm: FC<InputFormProps> = ({handleMessages}) => {
  const formRef = useRef<HTMLFormElement>(null)
  const handleSubmit: FormEventHandler = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const form = new FormData(e.currentTarget)
    const text = form.get('inputText')

    if (text && typeof text === 'string') {
      formRef.current?.reset()

      handleMessages(text)
    }
  }

  return (
    <form className={styles.container} ref={formRef} onSubmit={handleSubmit}>
      <label className={styles.label}>
        <input className={styles.textBox} type="text" name='inputText' placeholder="こんにちは"/>
      </label>
      <input className={styles.submitBtn} type="submit" value="送信"/>
    </form>
  )
}

export default InputForm
