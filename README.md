---

# Golang MongoDB Aggregation Example

This is an example Go program demonstrating the usage of MongoDB aggregation pipelines with the official MongoDB Go driver.

## Overview

The program connects to a MongoDB database and executes various aggregation queries on the collections to retrieve useful insights and information.

## Prerequisites

Before running the program, ensure you have the following installed:

- Go (Golang)
- MongoDB
- MongoDB Go Driver

## Setup

1. Clone this repository to your local machine.
2. Ensure MongoDB is running.
3. Install the required dependencies by running:

   ```bash
   go mod tidy
   ```

4. Modify the MongoDB connection details in the `databases` package as per your MongoDB setup.
5. Run the program:

   ```bash
   go run main.go
   ```

## Examples

The program executes several MongoDB aggregation queries:

1. **Example 01**: How many users are active.

   This example counts the number of active users in the `users` collection.

2. **Example 02**: What is the average age of all users.

   This example calculates the average age of all users in the `users` collection.

3. **Example 03**: List out the 5 top most common favorite fruits among users.

   This example lists the top 5 most common favorite fruits among users in the `users` collection, based on the frequency of occurrence.

4. **Example 04**: Find the total numbers of males and females.

   This example calculates the total number of males and females in the `users` collection.

5. **Example 05**: Which country has the highest number of registered users.

   This example determines the country with the highest number of registered users in the `users` collection.

6. **Example 06**: List all unique colors present in the collection.

   This example lists all unique eye colors present in the `users` collection.

7. **Example 07**: What is the average number of tags per user.

   This example calculates the average number of tags per user in the `users` collection.

8. **Example 08**: Another approach for calculating the average number of tags per user.

   This example demonstrates an alternative approach to calculate the average number of tags per user.

9. **Example 09**: How many users have "enum" as one of their tags.

   This example counts the number of users who have "enum" as one of their tags in the `users` collection.

10. **Example 10**: What are the names and ages of users who are inactive and have 'velit' as a tag.

    This example retrieves the names and ages of users who are inactive and have 'velit' as a tag in the `users` collection.

11. **Example 11**: How many users have a phone number starting with '+1(940)'.

    This example counts the number of users with a phone number starting with '+1(940)' in the `users` collection.

12. **Example 12**: Who has registered the most recently.

    This example retrieves the user who has registered most recently in the `users` collection.

13. **Example 13**: Categorize users by their favorite fruits.

    This example categorizes users by their favorite fruits in the `users` collection.

14. **Example 14**: How many users have 'ad' as the second tag in their list of tags.

    This example counts the number of users who have 'ad' as the second tag in their list of tags in the `users` collection.

15. **Example 15**: Find users who have both 'enum' and 'id' as their tags.

    This example retrieves users who have both 'enum' and 'id' as their tags in the `users` collection.

16. **Example 16**: List all companies located in the USA with their corresponding user count.

    This example lists all companies located in the USA along with their corresponding user count in the `users` collection.

17. **Example 17**: Retrieve author details for the book "The Great Gatsby".

    This example retrieves author details for the book "The Great Gatsby" from the `books` and `authors` collections.
