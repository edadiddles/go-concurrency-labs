# Lab 10 — Readers–Writers Lock

## Concept

Coordinated access with asymmetric concurrency.

This lab explores the readers–writers locking problem, where multiple readers may access a shared resource concurrently, but writers require exclusive access. Correctness depends on enforcing mutual exclusion, preventing starvation, and handling failures during access.

---

## Problem Statement

Implement a readers–writers lock that coordinates concurrent access to a shared resource.

Multiple reader workers may access the resource simultaneously. Writer workers require exclusive access and must block until all readers have exited and no other writers are active.

The system must ensure correctness, fairness, and clean termination under contention.

---

## Inputs

- An integer `R` representing the number of reader workers
- An integer `W` representing the number of writer workers
- A workload range controlling how long read and write operations take

Inputs may be hard-coded or provided via command-line arguments.

---

## Reader Behavior

Each reader must:

1. Begin execution
2. Request read access to the shared resource
3. Access the resource
4. Release read access
5. Repeat as configured
6. Terminate execution

Readers must not:
- Block other readers
- Access the resource while a writer holds exclusive access

---

## Writer Behavior

Each writer must:

1. Begin execution
2. Request exclusive write access
3. Access the shared resource
4. Release write access
5. Repeat as configured
6. Terminate execution

Writers must not:
- Access the resource concurrently with readers or other writers
- Starve indefinitely

---

## Output Requirements

- The program must report:
  - Total read operations
  - Total write operations
- Optional per-reader and per-writer access records may be produced
- Output ordering is not required to be deterministic

---

## Concurrency Constraints

- Multiple readers may access the resource concurrently
- Writers must have exclusive access
- No reader may access the resource while a writer is active
- Writers must not starve under continuous read load
- The main execution context must not exit until all readers and writers have completed

---

## Failure Injection (Optional)

The implementation should tolerate the following adversarial conditions:

- Continuous stream of readers
- Bursty writers
- Writer panic during access
- Reduced CPU availability

---

## Observations to Make

While running the program, observe:

- Whether readers block writers indefinitely
- Whether writers delay readers excessively
- Effects of writer failure
- Fairness under mixed workloads

---

## Success Criteria

The lab is considered complete when:

- Mutual exclusion guarantees are upheld
- Writers eventually acquire exclusive access
- Readers are allowed concurrent access when safe
- No deadlocks occur
- The program terminates deterministically

---

## Notes

This lab highlights fairness, starvation prevention, and failure handling in coordination primitives.

It concludes the coordination-focused labs and prepares the ground for cancellation and supervision.
