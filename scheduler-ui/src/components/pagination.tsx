import * as Primitive from "@/components/ui/pagination"
import { usePagination } from "@/hooks/use-pagination"

type PaginationProps = {
  totalItems: number
  defaultPage?: number
  defaultPageSize?: number
  siblingCount?: number
}

function getPageRange(
  currentPage: number,
  totalPages: number,
  siblingCount: number,
): (number | "ellipsis")[] {
  const totalNumbers = siblingCount * 2 + 5
  if (totalPages <= totalNumbers) {
    return Array.from({ length: totalPages }, (_, i) => i + 1)
  }

  const leftSibling = Math.max(currentPage - siblingCount, 1)
  const rightSibling = Math.min(currentPage + siblingCount, totalPages)

  const showLeftEllipsis = leftSibling > 2
  const showRightEllipsis = rightSibling < totalPages - 1

  const pages: (number | "ellipsis")[] = [1]

  if (showLeftEllipsis) {
    pages.push("ellipsis")
  } else {
    for (let i = 2; i < leftSibling; i++) pages.push(i)
  }

  for (let i = leftSibling; i <= rightSibling; i++) {
    if (i !== 1 && i !== totalPages) pages.push(i)
  }

  if (showRightEllipsis) {
    pages.push("ellipsis")
  } else {
    for (let i = rightSibling + 1; i < totalPages; i++) pages.push(i)
  }

  pages.push(totalPages)
  return pages
}

export default function Pagination({
  totalItems,
  defaultPage = 1,
  defaultPageSize = 10,
  siblingCount = 1,
}: PaginationProps) {
  const { currentPage, pageSize, nextPage, prevPage, goToPage } =
    usePagination({
      page: defaultPage,
      pageSize: defaultPageSize,
    })

  const totalPages = Math.max(1, Math.ceil(totalItems / pageSize))

  if (totalPages <= 1) return null

  const pages = getPageRange(currentPage, totalPages, siblingCount)
  const isFirst = currentPage <= 1
  const isLast = currentPage >= totalPages

  return (
    <Primitive.Pagination>
      <Primitive.PaginationContent>
        <Primitive.PaginationItem>
          <Primitive.PaginationPrevious
            href="#"
            aria-disabled={isFirst}
            data-disabled={isFirst}
            className={
              isFirst ? "pointer-events-none opacity-50" : undefined
            }
            onClick={(e) => {
              e.preventDefault()
              if (!isFirst) prevPage()
            }}
          />
        </Primitive.PaginationItem>

        {pages.map((page, i) =>
          page === "ellipsis" ? (
            <Primitive.PaginationItem key={`ellipsis-${i}`}>
              <Primitive.PaginationEllipsis />
            </Primitive.PaginationItem>
          ) : (
            <Primitive.PaginationItem key={page}>
              <Primitive.PaginationLink
                href="#"
                isActive={page === currentPage}
                onClick={(e) => {
                  e.preventDefault()
                  goToPage(page)
                }}
              >
                {page}
              </Primitive.PaginationLink>
            </Primitive.PaginationItem>
          ),
        )}

        <Primitive.PaginationItem>
          <Primitive.PaginationNext
            href="#"
            aria-disabled={isLast}
            data-disabled={isLast}
            className={isLast ? "pointer-events-none opacity-50" : undefined}
            onClick={(e) => {
              e.preventDefault()
              if (!isLast) nextPage()
            }}
          />
        </Primitive.PaginationItem>
      </Primitive.PaginationContent>
    </Primitive.Pagination>
  )
}
