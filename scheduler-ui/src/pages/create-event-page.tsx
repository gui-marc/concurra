import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Spinner } from "@/components/ui/spinner"
import { useCreateEvent } from "@/hooks/use-create-event"
import { Controller, useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"

import { z } from "zod"
import {
  Field,
  FieldDescription,
  FieldError,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { DatePickerWithRange } from "@/components/date-picker-with-range"
import React from "react"
import type { DateRange } from "react-day-picker"
import { Link } from "react-router"
import { ToggleGroup, ToggleGroupItem } from "@/components/ui/toggle-group"
import { Separator } from "@/components/ui/separator"

const formSchema = z.object({
  name: z.string().min(1, "Event name is required"),
  startTime: z.date(),
  endTime: z.date(),
  concurrencyTarget: z.number().min(1, "Concurrency target must be at least 1"),
})

type FormValues = z.infer<typeof formSchema>

const now = Date.now()

export default function CreateEventPage() {
  const [range, setRange] = React.useState<DateRange>({
    from: new Date(now),
    to: new Date(now + 60 * 60 * 1000 * 24), // Default to 1 day later
  })

  const { mutate, isPending } = useCreateEvent()

  const form = useForm<FormValues>({
    defaultValues: {
      name: "",
      startTime: new Date(now),
      endTime: new Date(now + 60 * 60 * 1000 * 24), // Default to 1 day later
      concurrencyTarget: undefined,
    },
    resolver: zodResolver(formSchema),
  })

  function onSubmit(data: FormValues) {
    mutate(data)
  }

  return (
    <div className="flex h-full w-full flex-1 items-center justify-center">
      <Card className="w-full max-w-lg">
        <CardHeader>
          <CardTitle>Schedule Event</CardTitle>
          <CardDescription>
            Fill in the details to create a new event.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <FieldGroup className="flex flex-row items-center justify-between gap-4">
              <div className="space-y-4">
                <Controller
                  name="name"
                  control={form.control}
                  render={({ field, fieldState }) => (
                    <Field data-invalid={!!fieldState.invalid}>
                      <FieldLabel>Name</FieldLabel>
                      <Input
                        {...field}
                        id="event-name"
                        aria-invalid={fieldState.invalid}
                        placeholder="Enter event name"
                        autoComplete="off"
                      />
                      {fieldState.invalid && (
                        <FieldError>{fieldState.error?.message}</FieldError>
                      )}
                    </Field>
                  )}
                />
                <Field data-invalid={!!form.formState.errors.startTime}>
                  <FieldLabel>Event Duration</FieldLabel>
                  <DatePickerWithRange
                    setRange={(range) => {
                      if (!range || !range.from || !range.to) return

                      setRange(range)

                      form.setValue("startTime", range.from)
                      form.setValue("endTime", range.to)
                    }}
                    range={range}
                  />
                  {form.formState.errors.startTime && (
                    <FieldError>
                      {form.formState.errors.startTime.message}
                    </FieldError>
                  )}
                </Field>
              </div>
              <Separator orientation="vertical" />
              <Field className="max-w-55">
                <FieldLabel>Concurrency Target</FieldLabel>
                <ToggleGroup
                  type="single"
                  {...form.register("concurrencyTarget")}
                  variant="outline"
                  spacing={2}
                  size="lg"
                >
                  <ToggleGroupItem
                    value="10"
                    aria-label="10"
                    className="flex size-16 flex-col items-center justify-center rounded-xl"
                    onClick={() => form.setValue("concurrencyTarget", 10)}
                  >
                    <span className="text-2xl leading-none font-bold">10</span>
                    <span className="text-xs text-muted-foreground">Small</span>
                  </ToggleGroupItem>
                  <ToggleGroupItem
                    value="15"
                    aria-label="15"
                    className="flex size-16 flex-col items-center justify-center rounded-xl"
                    onClick={() => form.setValue("concurrencyTarget", 15)}
                  >
                    <span className="text-2xl leading-none font-bold">15</span>
                    <span className="text-xs text-muted-foreground">
                      Medium
                    </span>
                  </ToggleGroupItem>
                  <ToggleGroupItem
                    value="30"
                    aria-label="30"
                    className="flex size-16 flex-col items-center justify-center rounded-xl"
                    onClick={() => form.setValue("concurrencyTarget", 30)}
                  >
                    <span className="text-2xl leading-none font-bold">30</span>
                    <span className="text-xs text-muted-foreground">Large</span>
                  </ToggleGroupItem>
                </ToggleGroup>
                {form.formState.errors.concurrencyTarget && (
                  <FieldError>
                    {form.formState.errors.concurrencyTarget.message}
                  </FieldError>
                )}
                <FieldDescription>
                  In{" "}
                  <code className="rounded-md bg-muted px-1 py-0.5 font-mono">
                    millions
                  </code>{" "}
                  of concurrent users.
                </FieldDescription>
              </Field>
            </FieldGroup>
          </form>
        </CardContent>
        <CardFooter className="flex justify-end gap-3">
          <Button asChild variant="outline">
            <Link to="/events">Cancel</Link>
          </Button>
          <Button onClick={form.handleSubmit(onSubmit)} disabled={isPending}>
            Schedule Event
            {isPending && <Spinner />}
          </Button>
        </CardFooter>
      </Card>
    </div>
  )
}
