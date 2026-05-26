export function formatDate(date: Date) {
  const formatter = Intl.DateTimeFormat("en-US", {
    dateStyle: "medium",
    timeStyle: "short",
  })

  return formatter.format(date)
}
