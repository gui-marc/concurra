import { create } from "axios"

export const client = create({
  baseURL: import.meta.env.VITE_API_BASE_URL || "http://localhost:8080",
  headers: {
    "Content-Type": "application/json",
  },
})

export type PageParams = {
  page: number
  pageSize: number
}

export type Page<T> = {
  items: T[]
  totalItems: number
}
