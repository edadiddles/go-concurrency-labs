# Lab 11 — Cooperative Cancellation

## Concept

Explicit cancellation and cooperative shutdown.

This lab explores cooperative cancellation in concurrent systems. Workers must detect cancellation signals, stop work promptly, and clean up safely. Correctness depends on responsiveness, cleanup guarantees, and ensuring cancellation does not leave the system in an inconsistent state.

---

## Problem Statement

Implement a system where concurrent workers perform ongoing work that may be canceled.

A cancellation signal may be issued at any time. Workers must observe the signal, terminate cooperatively, and release any resources they hold. The system must ensure clean shutdown without leaks or partial results.

---

## Inputs

- An integer `N` representing the number of workers
- A workload range controlling how long each unit of work takes
- A cancellation trigger determining when cancellation occurs

Inputs may be hard-coded or provided via command-line arguments.

---

## Worker Behavior

Each worker must:

1. Begin execution
2. Perform work in cancellable units
3. Periodically check for cancellation
4. Stop work promptly when cancellation is observed
5. Perform required cleanup
6. Terminate execution

Workers must not:
- Ignore cancellation signals
- Leave shared resources in an inconsistent state
- Block indefinitely after cancellation

---

## Cancellation Behavior

The cancellation mechanism must:

- Be observable by all workers
- Be idempotent
- Allow multiple cancellation signals
- Propagate promptly to all active workers

---

## Output Requirements

- The program must report:
  - Number of workers started
  - Number of workers canceled
  - Number of workers that completed normally
- Optional per-worker termination records may be produced
- Output ordering is not required to be deterministic

---

## Concurrency Constraints

- Workers must respond to cancellation within a bounded time
- No worker may continue processing after observing cancellation
- Cleanup must complete before program termination
- The main execution context must not exit until all workers have terminated

---

## Failure Injection (Optional)

The implementation should tolerate the following adversarial conditions:

- Cancellation during a critical section
- Cancellation during a blocking operation
- Repeated cancellation signals
- Reduced CPU availability

---

## Observations to Make

While running the program, observe:

- Cancellation responsiveness
- Cleanup correctness
- Effects of cancellation timing
- Consistency of final state

---

## Success Criteria

The lab is considered complete when:

- Cancellation stops work promptly
- All workers terminate cleanly
- No goroutine leaks occur
- Cleanup is performed correctly
- The program terminates deterministically

---

## Notes

This lab emphasizes cooperative design: cancellation is a request, not a force.

It prepares the groundwork for timeouts, supervisors, and fault-tolerant concurrency.
