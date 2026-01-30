# Lab 03 — Shared Counter

## Concept

Shared mutable state under concurrent access.

This lab introduces the simplest possible shared data structure—a counter—and explores how concurrent access without proper synchronization leads to lost updates, non-determinism, and correctness failures.

---

## Problem Statement

Implement a program that spawns a fixed number of concurrent workers, all of which increment a shared counter.

Each worker executes independently and repeatedly updates the same shared counter. The program must collect execution results and report the final counter value after all workers have completed.

Workers communicate only through shared memory.

---

## Inputs

- An integer `N` representing the number of workers
- An integer `K` representing the number of increments performed by each worker

Inputs may be hard-coded or provided via command-line arguments.

---

## Worker Behavior

Each worker must:

1. Begin execution
2. Perform `K` increments of a shared counter
3. Complete execution and report that it has finished

Workers must:
- Access the same shared counter
- Perform increments independently of other workers

Workers must not:
- Coordinate increment operations with other workers
- Assume atomicity of counter updates

---

## Output Requirements

- The program must report:
  - The expected counter value (`N × K`)
  - The observed final counter value
- Optional per-worker completion records may be produced
- Output ordering is not required to be deterministic

---

## Concurrency Constraints

- Exactly `N` workers must be created
- Workers must execute concurrently
- All workers must attempt to update the shared counter
- The main execution context must not exit until all workers have completed

---

## Failure Injection (Optional)

The implementation should tolerate the following adversarial conditions:

- Forced context switches during increment operations
- Random delays inside increment loops
- Extremely high contention (large `N`, large `K`)
- Reduced CPU availability

---

## Observations to Make

While running the program, observe:

- Whether the final counter value matches the expected value
- Frequency and magnitude of lost updates
- Run-to-run variability under identical inputs
- Sensitivity to scheduling pressure

---

## Success Criteria

The lab is considered complete when:

- The program consistently demonstrates incorrect results under concurrency
- Lost updates are observable without artificial instrumentation
- Repeated runs produce non-deterministic outcomes
- All workers complete execution
- The program terminates deterministically

---

## Notes

This lab intentionally introduces a data race. Correctness is *not* the goal.

This lab exists to make concurrency bugs observable, measurable, and undeniable. Synchronization mechanisms are intentionally deferred to later labs.
