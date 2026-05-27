import { Button } from "@/components/ui/button"
import {
  Empty,
  EmptyDescription,
  EmptyHeader,
  EmptyMedia,
  EmptyTitle,
} from "@/components/ui/empty"
import { Skeleton } from "@/components/ui/skeleton"
import EventsDataTable from "@/components/events/events-data-table"
import { useGetEvents } from "@/hooks/use-get-events"
import { ArrowUpRightIcon, HeartCrackIcon, PlusIcon } from "lucide-react"
import type React from "react"
import { Link, useSearchParams } from "react-router"
import Pagination from "@/components/pagination"

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

const PAGE_SIZE = 10

export default function EventsPage() {
  const [searchParams] = useSearchParams()
  const currentPage = Number(searchParams.get("page") || "1")

  const {
    data: page,
    isPending,
    isError,
  } = useGetEvents({
    page: currentPage,
    pageSize: PAGE_SIZE,
  })

  if (isError) return <Error />

  if (isPending || !page) return <Loading />

  const events = page.items || []

  return (
    <Page>
      <EventsDataTable events={events} />
      <div className="mt-6 flex items-center justify-between">
        <p className="text-sm text-muted-foreground">
          Showing {events.length} of {page.totalItems} events
        </p>
        <div>
          <Pagination
            totalItems={page.totalItems}
            defaultPageSize={PAGE_SIZE}
          />
        </div>
      </div>
    </Page>
  )
}
