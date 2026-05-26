import type { Event } from "@/api/events"
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb"
import {
  Card,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Skeleton } from "@/components/ui/skeleton"
import { formatDate } from "@/lib/date"
import type { eventLoader } from "@/loaders/events"
import { Link, useLoaderData } from "react-router"

export default function EventPage() {
  const event = useLoaderData<typeof eventLoader>() as Event | undefined

  if (!event) {
    return <Skeleton />
  }

  return (
    <div>
      <Header event={event} />
    </div>
  )
}

function Header({ event }: { event: Event }) {
  return (
    <header className="space-y-4">
      <Breadcrumb>
        <BreadcrumbList>
          <BreadcrumbItem>
            <BreadcrumbLink asChild>
              <Link to="/events">Events</Link>
            </BreadcrumbLink>
          </BreadcrumbItem>
          <BreadcrumbSeparator />
          <BreadcrumbItem>
            <BreadcrumbPage>{event.name}</BreadcrumbPage>
          </BreadcrumbItem>
        </BreadcrumbList>
      </Breadcrumb>
      <div className="grid grid-cols-1 gap-4 xl:grid-cols-2 2xl:grid-cols-4">
        <Card>
          <CardHeader>
            <CardDescription>Concurrency</CardDescription>
            <CardTitle className="text-2xl font-semibold tabular-nums @[250px]/card:text-3xl">
              {event.concurrencyTarget}{" "}
              <span className="text-sm font-normal text-muted-foreground">
                million users
              </span>
            </CardTitle>
          </CardHeader>
        </Card>
        <Card>
          <CardHeader>
            <CardDescription>Start Date</CardDescription>
            <CardTitle>{formatDate(new Date(event.startTime))}</CardTitle>
          </CardHeader>
        </Card>
        <Card>
          <CardHeader>
            <CardDescription>End Date</CardDescription>
            <CardTitle>{formatDate(new Date(event.endTime))}</CardTitle>
          </CardHeader>
        </Card>
      </div>
    </header>
  )
}
