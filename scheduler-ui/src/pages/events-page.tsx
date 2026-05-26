import { Button } from "@/components/ui/button"
import {
  Empty,
  EmptyDescription,
  EmptyHeader,
  EmptyMedia,
  EmptyTitle,
} from "@/components/ui/empty"
import {
  Item,
  ItemContent,
  ItemDescription,
  ItemTitle,
} from "@/components/ui/item"
import { Skeleton } from "@/components/ui/skeleton"
import { useGetEvents } from "@/hooks/use-get-events"
import { ArrowUpRightIcon, HeartCrackIcon, PlusIcon } from "lucide-react"
import type React from "react"
import { Link } from "react-router"

function Header() {
  return (
    <header className="mb-8 flex items-center justify-between">
      <div>
        <h1 className="text-lg font-bold text-foreground">Events</h1>
        <p className="text-sm text-muted-foreground">
          Manage your events and schedule
        </p>
      </div>
      <Button asChild>
        <Link to="/events/create">
          <PlusIcon /> Schedule Event
        </Link>
      </Button>
    </header>
  )
}

function Page({ children }: { children: React.ReactNode }) {
  return (
    <main>
      <Header />
      {children}
    </main>
  )
}

function Loading() {
  return (
    <Page>
      <ol className="space-y-4">
        {Array.from({ length: 5 }).map((_, i) => (
          <Skeleton className="h-10 w-full" key={i} />
        ))}
      </ol>
    </Page>
  )
}

function Error() {
  return (
    <Page>
      <Empty className="border border-dashed">
        <EmptyHeader>
          <EmptyMedia variant="icon">
            <HeartCrackIcon />
          </EmptyMedia>
          <EmptyTitle>Failed to load events</EmptyTitle>
          <EmptyDescription>
            An error occurred while fetching the events. Please try again later.
          </EmptyDescription>
        </EmptyHeader>

        <Button
          variant="link"
          asChild
          className="text-muted-foreground"
          size="sm"
        >
          <a target="_blank" href="https://github.com/gui-marc/concurra">
            Check Docs <ArrowUpRightIcon />
          </a>
        </Button>
      </Empty>
    </Page>
  )
}

export default function EventsPage() {
  const { data: events, isPending, isError } = useGetEvents()

  if (isError) return <Error />

  if (isPending || !events) return <Loading />

  return (
    <Page>
      <ol className="flex flex-col gap-6">
        {events.map((event) => (
          <Item key={event.id} variant="outline" asChild>
            <Link to={`/events/${event.id}`}>
              <ItemContent>
                <ItemTitle>{event.name}</ItemTitle>
                <ItemDescription>
                  \{new Date(event.startTime).toLocaleString()} -{" "}
                  {new Date(event.endTime).toLocaleString()}
                </ItemDescription>
              </ItemContent>
            </Link>
          </Item>
        ))}
      </ol>
    </Page>
  )
}
