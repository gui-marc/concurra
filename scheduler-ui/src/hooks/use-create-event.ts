import { createEvent } from "@/api/events"
import { useMutation } from "@tanstack/react-query"

export function useCreateEvent() {
  const mutation = useMutation({
    mutationKey: ["events"],
    mutationFn: createEvent,
  })

  return mutation
}
