import { client } from "./client"

const BASE_URL = "/api/v1/events"

export type Event = {
  id: string
  name: string
  startTime: Date
  endTime: Date
  concurrencyTarget: number
}

export async function getEvents(): Promise<Event[]> {
  const response = await client.get<Event[]>(BASE_URL)
  return response.data
}

export async function getEvent({ id }: { id: string }): Promise<Event> {
  const response = await client.get<Event>(`${BASE_URL}/${id}`)
  return response.data
}

export type CreateEventInput = Omit<Event, "id">

export async function createEvent(event: CreateEventInput): Promise<Event> {
  const response = await client.post<Event>(BASE_URL, event)
  return response.data
}

export async function updateEvent(event: Event): Promise<Event> {
  const response = await client.put<Event>(`${BASE_URL}/${event.id}`, event)
  return response.data
}

export async function deleteEvent({ id }: { id: string }): Promise<void> {
  await client.delete(`${BASE_URL}/${id}`)
}
