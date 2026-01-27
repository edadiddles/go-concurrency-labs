# Lab 01 — Parallel Workers

## Concept

Independent concurrent execution.

This lab explores the basic mechanics of launching multiple concurrent workers, observing their execution timing, and coordinating their completion without any shared mutable state.

---

## Problem Statement

Implement a program that spawns a fixed number of concurrent workers. Each worker executes independently, records timing information about its execution, and reports its results. The main execution context must wait for all workers to complete before exiting.

Workers do not communicate with each other and do not share memory.

---

## Inputs

- An integer `N` representing the number of workers to spawn
- A duration range `[min, max]` representing how long each worker should block

Inputs may be hard-coded or provided via command-line arguments.

---

## Worker Behavior

Each worker must:

1. Record the time at which it begins execution
2. Perform a blocking operation for a duration selected within the given range
3. Record the time at which it completes execution
4. Produce a single, complete output record containing:
   - Worker identifier
   - Start time
   - End time
   - Elapsed duration

Workers must not:
- Access shared mutable state
- Depend on the execution order of other workers

---

## Output Requirements

- Output must contain exactly one record per worker
- Each record must be complete and readable (no interleaved or partial output)
- Output ordering is not required to match worker creation order

---

## Concurrency Constraints

- Exactly `N` workers must be created
- Workers must execute concurrently
- The main execution context must not exit until all workers have completed
- No worker may block indefinitely

---

## Shutdown Semantics

- Program termination must be deterministic
- All resources must be released before exit
- No goroutines may remain running after program completion

---

## Failure Injection (Optional)

The implementation should tolerate the following adversarial conditions:

- Workers starting at different times
- Workers running for significantly different durations
- One or more workers completing almost immediately
- Artificial delays introduced in the main execution context before waiting

---

## Success Criteria

The lab is considered complete when:

- All workers execute independently and concurrently
- Exactly `N` output records are produced
- The program terminates only after all workers have completed
- Output is consistent and readable across multiple runs
- No goroutine leaks are observed

---

## Notes

This lab intentionally avoids shared state, message passing, or synchronization beyond completion tracking. It serves as a baseline for all subsequent concurrency labs.
