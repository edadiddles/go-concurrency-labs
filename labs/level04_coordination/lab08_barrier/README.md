# Lab 08 — Barrier

## Concept

Explicit synchronization and phase coordination.

This lab explores barrier synchronization, where a group of concurrent goroutines must all reach a common point before any may proceed. Correctness depends on precise coordination, correct reuse semantics, and robust handling of missing or delayed participants.

---

## Problem Statement

Implement a reusable barrier that coordinates a fixed number of concurrent workers.

Each worker performs some work, waits at a barrier, and then proceeds only after all workers have arrived. The barrier must release all waiting workers simultaneously once the required number of arrivals is reached.

The barrier may be used for one or more phases of execution.

---

## Inputs

- An integer `N` representing the number of workers
- An integer `P` representing the number of phases the barrier is used
- A workload range controlling how long each worker runs before reaching the barrier

Inputs may be hard-coded or provided via command-line arguments.

---

## Worker Behavior

Each worker must:

1. Begin execution
2. Perform phase-specific work
3. Arrive at the barrier
4. Wait until all `N` workers have arrived
5. Proceed to the next phase
6. Repeat for `P` phases
7. Terminate execution

Workers must not:
- Proceed past the barrier early
- Assume uniform arrival times

---

## Barrier Behavior

The barrier must:

- Block arriving workers until exactly `N` workers have arrived
- Release all waiting workers simultaneously
- Support reuse across multiple phases
- Prevent cross-phase interference

---

## Output Requirements

- The program must report:
  - Entry and exit times for each worker per phase
- Optional per-phase summary records may be produced
- Output ordering is not required to be deterministic

---

## Concurrency Constraints

- No worker may pass the barrier before all workers arrive
- No worker may be permanently blocked under normal operation
- Barrier reuse must not mix arrivals from different phases
- The main execution context must not exit until all workers complete all phases

---

## Failure Injection (Optional)

The implementation should tolerate the following adversarial conditions:

- One worker never arrives at the barrier
- Uneven arrival times across workers
- Barrier reused incorrectly
- Reduced CPU availability

---

## Observations to Make

While running the program, observe:

- Whether workers are released simultaneously
- Whether late arrivals delay the entire group
- Effects of missing arrivals
- Correctness of barrier reuse

---

## Success Criteria

The lab is considered complete when:

- All workers synchronize correctly at each phase
- No worker proceeds early
- No deadlocks occur under correct usage
- Barrier reuse behaves deterministically
- The program terminates cleanly

---

## Notes

This lab introduces strict phase-based coordination and highlights how small synchronization mistakes can lead to permanent blocking.

It sets the stage for one-time initialization and coordination primitives in subsequent labs.
