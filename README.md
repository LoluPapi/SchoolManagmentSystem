# School-Management-System
In this project, we are building a simple School management system, the school management system will allow school administrators to manage some activities and processes of the school. The system will provides features such as

- Student Management
- Lectures/Courses Management


We Developed a RESTFUL api to handle the requirements of this system using Go(Gin) and MySql database. 
Your Api should be fluent,and follow the best practices. Your api should provide fluent CRUD endpoints for the entities (Students, Courses) and also provide endpoints for the following features

- List courses taken by a student 
- List students taking a particular course
- Update list of courses taken by a student


**Usage**

To build and run the School Management System API, please follow the instructions below:

- Clone the GitHub repository to your local machine.
- Install Docker and Docker Compose.
- Open a terminal in the project's root directory and run docker-compose up -d. This will start the MySQL database and the API server in the background.
- To stop the containers, run docker-compose down.
- To access the OpenAPI documentation, go to http://localhost:8080/swagger/index.html in your web browser. This will display the Swagger UI for the API, where you can see the available endpoints and test them out.
