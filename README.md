Here’s an updated `README.md` for Windows-only instructions:

---

# **Go Template CLI**

A CLI tool to generate new Go projects with a predefined template structure. Simplify your project initialization process and ensure consistency in folder and file organization.

---

## **Features**
- Quickly generate a new Go project based on a standardized template.
- Predefined folder structure for organized codebases.
- Lightweight and easy to use.

---

## **Folder Structure**

Below is the default folder structure that will be generated:

```
project-name/
├── cmd/
│   └── main.go
├── internal/
├── pkg/
└── go.mod
```

---

## **Usage**

### **1. Setup the CLI**

#### **Build the CLI**
1. Navigate to the `go-template-cli` folder containing `main.go`.
2. Run the following command in your terminal to build the executable:
   ```bash
   go build -o go-template-cli.exe main.go
   ```

#### **Add the Executable to PATH**
To make the CLI accessible from anywhere:
1. Locate the folder containing `go-template-cli.exe`.
2. Open the **Start Menu** and search for "Environment Variables."
3. Click **Edit the system environment variables**.
4. In the **System Properties** window, click **Environment Variables**.
5. Under **System variables**, find and select the `Path` variable, then click **Edit**.
6. Click **New** and add the full path to the folder containing `go-template-cli.exe`.
7. Click **OK** to save your changes.

### **2. Create a New Project**
Open a terminal and navigate to the directory where you want to create your project. Run the following command:
```bash
go-template-cli create <project-name>
```

This will:
- Create a new folder named `<project-name>`.
- Copy all contents of the `template/` folder into the newly created project.

---

## **Prerequisites**
- [Golang](https://golang.org/) (1.19 or later installed on your system).
- Basic knowledge of Golang and terminal usage.

---

## **Contributing**
Contributions are welcome! To contribute:
1. Fork the repository.
2. Create a new branch.
3. Make your changes.
4. Submit a pull request.

---

## **License**
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Let me know if you need additional customization!
