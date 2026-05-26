import type { PageParams } from "@/api/client"
import { useSearchParams } from "react-router"

export function usePagination(defaultParams: PageParams) {
  const [pageParams, setPageParams] = useSearchParams({
    page: String(defaultParams.page),
    pageSize: String(defaultParams.pageSize),
  })

  const currentPage =
    Number(pageParams.get("page")) || defaultParams.page
  const pageSize =
    Number(pageParams.get("pageSize")) || defaultParams.pageSize

  function setPage(page: number) {
    setPageParams(
      (prev) => {
        const next = new URLSearchParams(prev)
        next.set("page", String(page))
        if (!next.has("pageSize")) next.set("pageSize", String(pageSize))
        return next
      },
      { replace: true },
    )
  }

  function nextPage() {
    setPage(currentPage + 1)
  }

  function prevPage() {
    setPage(Math.max(currentPage - 1, 1))
  }

  function goToPage(page: number) {
    setPage(Math.max(page, 1))
  }

  return {
    pageParams,
    setPageParams,
    nextPage,
    goToPage,
    prevPage,
    currentPage,
    pageSize,
  }
}
