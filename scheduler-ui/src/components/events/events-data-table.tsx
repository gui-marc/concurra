import type { Event } from "@/api/events"
import { DataTable } from "../data-table"
import { columns } from "./events-table-columns"

export default function EventsTableColumns({ events }: { events: Event[] }) {
  return <DataTable columns={columns} data={events} />
}
