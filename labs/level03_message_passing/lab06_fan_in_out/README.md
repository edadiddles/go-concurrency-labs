# Lab 06 — Fan-Out / Fan-In

## Concept

Parallel execution with centralized result aggregation.

This lab explores the fan-out / fan-in concurrency pattern, where work is distributed to multiple workers in parallel and results are collected and merged into a single stream. Correctness depends on backpressure handling, shutdown ordering, and proper coordination between producers, workers, and collectors.

---

## Problem Statement

Implement a system that fans out work items to multiple concurrent workers and fans in their results to a single collector.

A producer submits work items. Workers process work items concurrently. A result collector receives and aggregates results from all workers.

The system must ensure that all work is processed, all results are collected exactly once, and all goroutines terminate cleanly.

---

## Inputs

- An integer `N` representing the number of work items
- An integer `W` representing the number of workers
- A workload range controlling how long each work item takes to execute

Inputs may be hard-coded or provided via command-line arguments.

---

## Producer Behavior

The producer must:

1. Begin execution
2. Submit `N` work items
3. Stop submitting work
4. Terminate execution

The producer must not:
- Submit duplicate work items
- Assume workers process work at uniform speed

---

## Worker Behavior

Each worker must:

1. Begin execution
2. Receive work items
3. Process each work item independently
4. Produce exactly one result per work item
5. Continue until no more work is available
6. Terminate execution

Workers must not:
- Coordinate with other workers
- Retain results after emission

---

## Collector Behavior

The collector must:

1. Begin execution
2. Receive results from all workers
3. Aggregate results
4. Detect completion
5. Terminate execution

The collector must not:
- Exit before all results are received
- Assume result ordering

---

## Output Requirements

- The program must report:
  - Total work items submitted
  - Total results collected
- Optional per-worker or per-result records may be produced
- Output ordering is not required to be deterministic

---

## Concurrency Constraints

- Each work item must produce exactly one result
- No result may be duplicated or lost
- Workers must execute concurrently
- The collector must not block indefinitely
- The main execution context must not exit until:
  - All work is processed
  - All results are collected
  - All goroutines have terminated

---

## Failure Injection (Optional)

The implementation should tolerate the following adversarial conditions:

- One worker stalling indefinitely
- The collector exiting early
- Input closing unexpectedly
- Mixed fast and slow workers

---

## Observations to Make

While running the program, observe:

- Backpressure behavior when the collector is slow
- Effects of uneven worker performance
- Shutdown correctness and ordering
- Whether stalled workers affect system liveness

---

## Success Criteria

The lab is considered complete when:

- All submitted work items produce exactly one result
- All results are collected
- No goroutine leaks occur
- The program terminates deterministically
- Behavior is stable across repeated runs

---

## Notes

This lab emphasizes result aggregation, backpressure propagation, and clean shutdown in multi-stage concurrent systems.

It prepares the groundwork for pipeline-based concurrency patterns.
