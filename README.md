# 📝 Todo App

A simple **Todo Application** built with **Go** and **HTML Templates**.  
Users can add new tasks, delete tasks, and mark them as **done/undone**.  
All data is stored in todos.json (no database).

---

## 🚀 Features
- ➕ Add new tasks  
- ❌ Delete tasks  
- ✅ Mark tasks as done / undone  


## 🛠️ Tech Stack
- **Go (Golang)**  
- **Go html/template**  



## 📦 Setup & Usage
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/todo-app.git
2. Navigate to the project folder:
   ```bash
    cd todo-app
3. Open in your browser:
   ```bash
    http://localhost:8080


## 🌐 Endpoints
| Method | Path           | Description                       |
| ------ | -------------- | --------------------------------- |
| `GET`  | `/`            | Show all todos                    |
| `GET`  | `/add`         | Show form to add a new todo       |
| `POST` | `/add`         | Add a new todo (from form submit) |
| `GET`  | `/delete/{id}` | Delete todo by ID                 |
| `GET`  | `/toggle/{id}` | Toggle done/undone by ID          |

## 👨‍💻 Author
**[Vahid Arzumanov](https://github.com/27vhd)**

---
