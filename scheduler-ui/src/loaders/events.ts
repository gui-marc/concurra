import { getEvent } from "@/api/events"
import { queryClient } from "@/lib/query-client"
import { redirect, type LoaderFunction } from "react-router"

export const eventLoader: LoaderFunction = async ({ params }) => {
  const { id } = params
  if (!id) {
    return redirect("/events")
  }

  const event = await queryClient.fetchQuery({
    queryKey: ["events", id],
    queryFn: () => {
      return getEvent({ id })
    },
  })

  return event
}
