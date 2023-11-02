# Database Diagram Documentation

## Overview

The diagram showcases a database structure for a pet management system, encompassing pets, owners, appointments, payments, and related services. There are five main tables: `pets`, `owners`, `payment`, `appointment`, and `services`, which have defined relationships amongst themselves.

![ED_Diagram](/docs/img/ER-Diagram.png)

## Tables and Relationships

### Table `pets`

**Description**: Stores information regarding pets.

**Columns**:

- `id_pet`: Unique identifier for each pet.
- `name_pet`: Name of the pet.
- `birth_date_pet`: Pet's birthdate.
- `visit_counter_pet`: Counter of visits the pet has had.
- `owners_pet`: Relationship to the `owners` table, possibly indicating the number of owners a pet has.

### Table `owners`

**Description**: Contains data related to pet owners.

**Columns**:

- `id_owner`: Unique identifier for each owner.
- `name_owner`: Name of the owner.
- `birth_date_owner`: Owner's birthdate.
- `phone_owner`: Owner's phone number.
- `pets_owner`: Relationship to the `pets` table, possibly indicating the number of pets an owner has.

### Table `payment`

**Description**: Retains information about the payments made for services.

**Columns**:

- `id_payment`: Unique identifier for each payment.
- `id_owner`: Link to the `owners` table indicating who made the payment.
- `id_appointment`: Link to the `appointment` table indicating which appointment the payment is for.
- `amount_payment`: Amount paid.
- `state_payment`: Status of the payment (e.g., pending, completed).

### Table `appointment`

**Description**: Stores information about scheduled appointments for pets.

**Columns**:

- `id_appointment`: Unique identifier for each appointment.
- `mascota_id`: Relationship to the `pets` table indicating which pet the appointment is for.
- `owner_id`: Relationship to the `owners` table indicating the owner of the pet that has the appointment.
- `date_appointment`: Date and time of the appointment.
- `next_date_appointment`: Date and time of the next scheduled appointment.
- `real_next_date_appointment`: Actual date and time of the next appointment.
- `services_appointment`: Relationship to the `services` table indicating which services will be rendered during the appointment.

### Table `services`

**Description**: Holds data on available services for pets.

**Columns**:

- `id_service`: Unique identifier for each service.
- `name_service`: Name of the service.
- `cost_service`: Cost of the service.

## Relationships

- `pets` has a relationship with `owners` via `pets_owner`, potentially indicating how many pets an owner has.
- `pets` has a relationship with `owners` via `owners_pet`, potentially indicating how owners a pet has.
- `payment` is linked to both `owners` (who makes the payment) and `appointment` (what appointment the payment is for).
- `appointment` has relationships with `pets` (which pet the appointment is for), `owners` (owner of the pet), and `services` (services provided in the appointment).
