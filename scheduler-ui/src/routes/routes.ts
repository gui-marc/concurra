import { createBrowserRouter } from "react-router"

import RootLayout from "@/layout/root-layout"
import EventPage from "@/pages/event-page"
import EventsPage from "@/pages/events-page"
import CreateEventPage from "@/pages/create-event-page"

export const router = createBrowserRouter([
  {
    path: "/",
    Component: RootLayout,
    children: [
      {
        index: true,
        path: "events",
        Component: EventsPage,
      },
      {
        path: "events/:event",
        Component: EventPage,
      },
      {
        path: "events/create",
        Component: CreateEventPage,
      },
    ],
  },
])
