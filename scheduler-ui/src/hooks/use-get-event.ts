import { getEvent } from "@/api/events"
import { useQuery } from "@tanstack/react-query"

export function useGetEvent(eventId: string) {
  const query = useQuery({
    queryKey: ["events", eventId],
    queryFn: () => getEvent({ id: eventId }),
  })

  return query
}
