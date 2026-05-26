import { Outlet } from "react-router"

export default function RootLayout() {
  return (
    <div className="mx-auto flex min-h-svh max-w-7xl flex-col p-4 pt-8 pb-10">
      <div className="flex h-full w-full flex-1 flex-col">
        <Outlet />
      </div>
    </div>
  )
}
