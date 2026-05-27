import type { Event } from "@/api/events"
import { DataTable } from "../data-table"
import { columns } from "./events-table-columns"
import {
  Empty,
  EmptyContent,
  EmptyDescription,
  EmptyHeader,
  EmptyMedia,
  EmptyTitle,
} from "../ui/empty"
import { FolderIcon } from "lucide-react"
import { Button } from "../ui/button"
import { Link } from "react-router"

export default function EventsTableColumns({ events }: { events: Event[] }) {
  return (
    <DataTable columns={columns} data={events}>
      <Empty>
        <EmptyHeader>
          <EmptyMedia variant="icon">
            <FolderIcon />
          </EmptyMedia>
          <EmptyTitle>No events found</EmptyTitle>
          <EmptyDescription>
            You haven't scheduled any events yet. Click the button below to
            create your first event.
          </EmptyDescription>
        </EmptyHeader>
        <EmptyContent>
          <Button asChild>
            <Link to="/events/create">Create Event</Link>
          </Button>
        </EmptyContent>
      </Empty>
    </DataTable>
  )
}
