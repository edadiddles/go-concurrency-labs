# Lab 09 — Run-Once Initialization

## Concept

Single-execution initialization under concurrent access.

This lab explores run-once initialization, where multiple concurrent goroutines may attempt to initialize a shared resource, but initialization must occur exactly once. Correctness depends on preventing partial initialization, double execution, and unsafe access during initialization.

---

## Problem Statement

Implement a run-once initialization mechanism for a shared resource accessed by multiple concurrent workers.

Workers may attempt to access the resource concurrently. If the resource has not yet been initialized, exactly one worker must perform initialization while all others wait. Once initialization completes, all workers must observe a fully initialized resource.

---

## Inputs

- An integer `N` representing the number of workers
- A workload range controlling how long initialization and worker access take

Inputs may be hard-coded or provided via command-line arguments.

---

## Worker Behavior

Each worker must:

1. Begin execution
2. Attempt to access the shared resource
3. Trigger initialization if the resource is uninitialized
4. Block or wait if initialization is in progress
5. Proceed once initialization completes
6. Use the initialized resource
7. Terminate execution

Workers must not:
- Perform initialization more than once
- Access the resource before initialization completes

---

## Initialization Behavior

The initialization logic must:

- Execute exactly once
- Be visible to all workers after completion
- Either complete successfully or fail deterministically
- Prevent partial initialization from being observed

---

## Output Requirements

- The program must report:
  - Whether initialization occurred
  - How many workers accessed the resource
- Optional per-worker access records may be produced
- Output ordering is not required to be deterministic

---

## Concurrency Constraints

- Initialization must occur at most once
- All workers must observe the same initialized state
- No worker may access the resource before initialization completes
- The main execution context must not exit until all workers complete

---

## Failure Injection (Optional)

The implementation should tolerate the following adversarial conditions:

- Initialization failure
- Slow initialization
- Multiple workers attempting initialization simultaneously
- Reduced CPU availability

---

## Observations to Make

While running the program, observe:

- Whether initialization executes more than once
- Whether workers block correctly during initialization
- Effects of initialization failure
- Visibility of initialized state across workers

---

## Success Criteria

The lab is considered complete when:

- Initialization executes exactly once
- No worker observes partial initialization
- All workers eventually proceed after initialization
- The program terminates deterministically
- Behavior is stable across repeated runs

---

## Notes

This lab emphasizes safe publication, memory visibility, and correctness under contention.

It prepares the ground for reader–writer coordination and more complex synchronization patterns.
