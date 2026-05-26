import { deleteEvent } from "@/api/events"
import { useMutation } from "@tanstack/react-query"

export function useDeleteEvent(eventId: string) {
  const mutation = useMutation({
    mutationKey: ["events", eventId],
    mutationFn: () => deleteEvent({ id: eventId }),
  })

  return mutation
}
