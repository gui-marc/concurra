import { getEvents } from "@/api/events"
import { useQuery } from "@tanstack/react-query"

export function useGetEvents() {
  const query = useQuery({
    queryKey: ["events"],
    queryFn: getEvents,
  })

  return query
}
