# Lab 12 — Timeout Queue

## Concept

Time-bounded work and deadline enforcement.

This lab explores how to enforce time limits on queued work. Tasks that wait too long before being processed must be canceled or discarded deterministically, without leaking resources or blocking the system.

---

## Problem Statement

Implement a work queue where each submitted task has an associated timeout.

Workers process tasks from the queue. If a task is not started before its timeout expires, it must be canceled and must not be processed. The system must handle timeouts cleanly while continuing to process valid work.

---

## Inputs

- An integer `N` representing the number of worker goroutines
- A queue capacity
- A task timeout duration
- A workload duration per task

Inputs may be hard-coded or provided via command-line arguments.

---

## Task Behavior

Each task must:

1. Be enqueued with an associated timeout
2. Either:
   - Begin execution before its timeout expires, or
   - Be canceled due to timeout
3. Produce exactly one terminal outcome:
   - Completed
   - Timed out

Tasks must not:
- Execute after timing out
- Produce multiple outcomes
- Block system shutdown

---

## Worker Behavior

Each worker must:

1. Fetch tasks from the queue
2. Respect task timeouts
3. Execute only valid (non-expired) tasks
4. Report task completion or cancellation
5. Terminate cleanly during shutdown

Workers must not:
- Execute expired tasks
- Ignore cancellation or timeout signals
- Leak goroutines when idle

---

## Timeout Semantics

The timeout mechanism must:

- Be enforced per task
- Be monotonic and race-free
- Prevent late execution
- Allow tasks to time out while waiting or during dispatch

---

## Output Requirements

The program must report:

- Number of tasks submitted
- Number of tasks completed
- Number of tasks timed out
- Number of tasks discarded without execution

Optional per-task records may be produced.

---

## Concurrency Constraints

- Queue access must be safe under concurrent producers and consumers
- Timeout handling must not block the queue
- Workers must remain responsive under load
- Shutdown must wait for all workers and timeout handlers

---

## Failure Injection (Optional)

The implementation should tolerate the following adversarial conditions:

- Tasks timing out immediately
- Workers running slower than expected
- Bursty task submission
- Reduced worker count
- Artificial delays in timeout enforcement

---

## Observations to Make

While running the program, observe:

- Accuracy of timeout enforcement
- Fairness of task execution
- System behavior under overload
- Interaction between timeouts and shutdown

---

## Success Criteria

The lab is considered complete when:

- No expired task is executed
- All tasks reach exactly one terminal state
- Workers remain responsive
- Shutdown is deterministic and clean
- No goroutine or timer leaks occur

---

## Notes

This lab introduces time as a correctness constraint.

Correct solutions require careful coordination between queues, workers, and cancellation mechanisms, and serve as a foundation for deadline-aware schedulers and real-time systems.
