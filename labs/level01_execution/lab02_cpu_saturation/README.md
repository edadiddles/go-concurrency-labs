# Lab 02 — Saturated Execution

## Concept

Independent concurrent execution under saturation.

This lab explores how independent workers behave when system resources—particularly CPU and scheduler capacity—become the limiting factor. No coordination is introduced between workers; instead, pressure is applied by oversubscribing execution resources.

---

## Problem Statement

Implement a program that spawns a fixed number of independent concurrent workers and drives the system into a saturated execution state.

Each worker executes independently, records timing information, and reports its results. The program must allow different classes of work to be executed (CPU-bound and blocking) in order to observe scheduling behavior under load.

Workers do not communicate with each other and do not share memory.

---

## Inputs

- An integer `N` representing the number of workers to spawn
- A mode indicating the type of work performed by workers:
  - CPU-bound
  - Blocking
  - Mixed
- A duration or workload range controlling how long each worker runs

Inputs may be hard-coded or provided via command-line arguments.

---

## Worker Behavior

Each worker must:

1. Record the time at which it begins execution
2. Perform a unit of work according to the selected mode:
   - CPU-bound computation
   - Blocking operation
3. Record the time at which it completes execution
4. Produce a single, complete output record containing:
   - Worker identifier
   - Start time
   - End time
   - Elapsed duration

Workers must not:
- Access shared mutable state
- Coordinate execution with other workers
- Assume availability of CPU resources

---

## Output Requirements

- Output must contain exactly one record per worker
- Each record must be complete and readable (no interleaved or partial output)
- Output ordering is not required to match worker creation order

---

## Concurrency Constraints

- Exactly `N` workers must be created
- Workers must be able to execute concurrently, subject to system limits
- Oversubscription of CPU resources must be permitted
- The main execution context must not exit until all workers have completed
- No worker may block indefinitely

---

## Saturation Scenarios

The implementation must allow the system to be driven into saturation through one or more of the following:

- Creating significantly more workers than available CPU cores
- Mixing CPU-bound and blocking workers
- Restricting available CPU resources (e.g., reduced parallelism)

---

## Failure Injection (Optional)

The implementation should tolerate the following adversarial conditions:

- Worker count greatly exceeding available CPU cores
- Mixed workloads with uneven execution characteristics
- Artificial system load applied during execution
- One or more workers exiting earlier than expected

---

## Observations to Make

While running the program, observe:

- Scheduling fairness among workers
- Variance in worker runtimes
- Evidence of starvation or delayed execution
- Changes in throughput as saturation increases

---

## Success Criteria

The lab is considered complete when:

- All non-failing workers complete execution
- Exactly `N` output records are produced
- Wall-clock runtime is bounded by:
 
`max(worker_duration) ≤ total_runtime ≤ sum(worker_durations)`

- Evidence of concurrent execution exists
- The program terminates deterministically
- No goroutine leaks are observed

---

## Notes

This lab intentionally avoids shared state, synchronization primitives, or message passing. It focuses exclusively on understanding how the Go runtime schedules independent work under pressure.

This lab serves as a bridge between basic concurrency mechanics and the introduction of shared state in subsequent labs.
