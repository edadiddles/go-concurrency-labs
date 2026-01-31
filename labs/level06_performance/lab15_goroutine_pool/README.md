# Lab 15 — Goroutine Pool

## Concept

Controlled concurrency and backpressure.

This lab explores how a fixed-size pool of worker goroutines can be used to limit concurrency, absorb workload bursts, and provide predictable resource usage. The focus is on task scheduling, lifecycle management, and safe shutdown under load and failure.

---

## Problem Statement

Implement a goroutine pool that accepts tasks, executes them using a bounded number of workers, and supports graceful shutdown.

The pool must prevent unbounded goroutine creation while ensuring that submitted tasks are either executed or explicitly rejected.

---

## Inputs

- Pool size `P` (maximum number of concurrent workers)
- Task submission rate or workload size
- Optional queue capacity
- Shutdown behavior configuration

Inputs may be hard-coded or provided via command-line arguments.

---

## Task Behavior

Each task must:

1. Execute independently
2. Optionally block or perform CPU-bound work
3. Optionally fail or panic
4. Produce a completion record

Tasks must not:
- Access shared mutable state without synchronization
- Assume execution order
- Assume immediate execution upon submission

---

## Pool Behavior

The goroutine pool must:

- Maintain at most `P` active worker goroutines
- Accept tasks via a defined submission mechanism
- Apply backpressure when overloaded
- Support orderly shutdown

---

## Output Requirements

The program must report:

- Number of tasks submitted
- Number of tasks executed
- Number of tasks rejected or dropped (if applicable)
- Worker lifecycle events (optional)

Output must remain readable and consistent under high load.

---

## Concurrency Constraints

- No unbounded goroutine creation
- No task may be executed more than once
- No task may be silently lost
- Pool shutdown must be deterministic

---

## Failure Injection (Optional)

The implementation should tolerate the following adversarial conditions:

- Burst submission of tasks
- Tasks that panic
- Tasks that block indefinitely
- Pool shutdown during active task execution

---

## Observations to Make

While running the program, observe:

- Backpressure behavior under burst load
- Task latency vs throughput tradeoffs
- Worker reuse effectiveness
- Behavior during shutdown

---

## Success Criteria

The lab is considered complete when:

- Concurrency is bounded as configured
- Backpressure is observable and measurable
- Task failures do not crash the pool
- Shutdown completes without leaks or deadlocks
- Output remains consistent across runs

---

## Notes

This lab ties together concepts from contention, message passing, and failure handling. Design tradeoffs between simplicity, fairness, and performance should be explicitly considered.
