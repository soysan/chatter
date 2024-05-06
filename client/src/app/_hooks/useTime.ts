import { useCallback } from 'react'

export const useTime = () => {
  const addZero = (num: number): string => {
    return num < 10 ? `0${num}` : "" + num
  }

  const displayTime = useCallback((datetime?: string): string => {
    const date = datetime ? new Date(datetime) : new Date()
    const hours = date.getHours()
    const minutes = date.getMinutes()
    const seconds = date.getSeconds()

    return `${addZero(hours)}:${addZero(minutes)}:${addZero(seconds)}`
  }, [])

  return {displayTime}
}