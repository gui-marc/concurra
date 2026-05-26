import type { PageParams } from "@/api/client"
import { getEvents } from "@/api/events"
import { useQuery } from "@tanstack/react-query"

export function useGetEvents({ page, pageSize }: PageParams) {
  const query = useQuery({
    queryKey: ["events", page, pageSize],
    queryFn: () => getEvents({ page, pageSize }),
  })

  return query
}
