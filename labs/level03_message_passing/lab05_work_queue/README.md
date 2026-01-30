# Lab 05 — Work Queue

## Concept

Coordination through message passing.

This lab replaces shared mutable state with explicit message passing. Work is distributed to a pool of workers via a queue, and results are collected asynchronously. Correctness depends on channel discipline, ownership, and lifecycle management.

---

## Problem Statement

Implement a concurrent work queue that distributes tasks from one or more producers to a pool of workers using message passing.

Producers submit work items into the queue. Workers process work items independently and produce results. The system must ensure that all submitted work is processed exactly once and that all goroutines terminate cleanly.

---

## Inputs

- An integer `P` representing the number of producers
- An integer `W` representing the number of workers
- An integer `N` representing the number of tasks each producer submits
- A workload range controlling how long each task takes to execute

Inputs may be hard-coded or provided via command-line arguments.

---

## Producer Behavior

Each producer must:

1. Begin execution
2. Submit `N` work items into the work queue
3. Stop submitting work after `N` items
4. Terminate execution

Producers must not:
- Submit duplicate work items
- Assume workers are available when submitting work

---

## Worker Behavior

Each worker must:

1. Begin execution
2. Receive work items from the queue
3. Process each work item independently
4. Produce exactly one result per work item
5. Terminate execution when no more work is available

Workers must not:
- Process the same work item more than once
- Retain ownership of work items after completion

---

## Output Requirements

- The program must report:
  - Total work items submitted
  - Total work items processed
  - Total results produced
- Optional per-worker result records may be produced
- Output ordering is not required to be deterministic

---

## Concurrency Constraints

- Each submitted work item must be processed exactly once
- No work item may be lost or duplicated
- Workers must operate concurrently
- The work queue must not deadlock
- The main execution context must not exit until all work is processed and all workers have terminated

---

## Failure Injection (Optional)

The implementation should tolerate the following adversarial conditions:

- One or more producers terminating early
- A worker panicking during task execution
- Slow workers mixed with fast workers
- Reduced CPU availability

---

## Observations to Make

While running the program, observe:

- Whether all work items are eventually processed
- Whether worker failures affect unrelated work
- Whether goroutines terminate cleanly
- Whether work distribution remains balanced

---

## Success Criteria

The lab is considered complete when:

- All submitted work items are processed exactly once
- All results are produced
- No goroutines remain blocked or leaked
- The program terminates deterministically
- Behavior is stable across repeated runs

---

## Notes

This lab emphasizes ownership transfer, channel lifecycle management, and explicit coordination via message passing.

It serves as the foundation for fan-out/fan-in and pipeline patterns in subsequent labs.
