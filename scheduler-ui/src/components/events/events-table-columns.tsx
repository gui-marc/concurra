import type { Event } from "@/api/events"
import { type ColumnDef } from "@tanstack/react-table"
import { Button } from "../ui/button"
import { Link } from "react-router"
import { ArrowRightIcon } from "lucide-react"

export const columns: ColumnDef<Event>[] = [
  {
    accessorKey: "name",
    header: "Name",
  },
  {
    accessorKey: "startTime",
    header: "Start Time",
    cell: ({ row }) => {
      const formatted = Intl.DateTimeFormat("en-US", {
        dateStyle: "medium",
        timeStyle: "short",
      }).format(new Date(row.original.startTime))
      return <span>{formatted}</span>
    },
  },
  {
    accessorKey: "endTime",
    header: "End Time",
    cell: ({ row }) => {
      const formatted = Intl.DateTimeFormat("en-US", {
        dateStyle: "medium",
        timeStyle: "short",
      }).format(new Date(row.original.endTime))
      return <span>{formatted}</span>
    },
  },
  {
    accessorKey: "concurrencyTarget",
    header: "Concurrency Target",
  },
  {
    id: "actions",
    cell: ({ row }) => {
      return (
        <div className="flex justify-end">
          <Button variant="link" size="icon" asChild>
            <Link to={`/events/${row.original.id}`}>
              <ArrowRightIcon />
            </Link>
          </Button>
        </div>
      )
    },
  },
]
