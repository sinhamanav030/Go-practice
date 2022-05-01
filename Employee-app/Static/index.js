// Function to load data on Opening Page
const getData = () =>{
    let Employees =[];

    if(window.sessionStorage.getItem('Employees')==undefined){
        window.sessionStorage.setItem('Employees',JSON.stringify(Employees));
    }else{
        Employees = JSON.parse(window.sessionStorage.getItem('Employees'));
    }
    // const table = document.getElementById('Employee-table');
    const table = $("#Employee-table").DataTable();

    Employees.forEach((emp,i)=>{
        table.row.add([emp.name,emp.code,emp.email,emp.gender,emp.designation,emp.dob,`<button class="btn" onclick="updateRow(${i})"><i class = "material-icons" style="color:blue">edit</i></button>`,`<button class="btn" onclick="deleteRow(${i})"><i class="material-icons" style="color:red;">delete</i></button>`]).draw()
    });
}

//function to delete employee
const deleteRow = (i)=>{
    // console.log(i);
    let Employees = JSON.parse(window.sessionStorage.getItem('Employees'));
    let up_employees =[]; 
    Employees.forEach((emp,id)=>{
        if(id != i){
            up_employees.push(emp);
        }
    });
    console.log(up_employees);
    window.sessionStorage.setItem('Employees',JSON.stringify(up_employees));
    window.location.pathname="./index.html";
}

//function to update employee
const updateRow =(i) =>{
    let Employees = JSON.parse(window.sessionStorage.getItem('Employees'));
    Employees.forEach((emp,id)=>{
        if(id == i){
            document.getElementById('update-code').setAttribute('value',emp.code);
            $('#updateModal').modal('show');
            return;
            // window.open("./update.html","MsgWindow","height=600,width=400,resizable=no,top=150px,left=600px");
            // window.location.pathname="./update.html";
        }
    });
}

//From validation on adding user
const addFormValidation = (form)=>{
    let codeRegEx = /^[a-zA-Z0-9]*$/;
    let emailRegEx = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    let nameRegex = /^[a-zA-Z ]*$/;
    let desRegEx = /^[a-zA-Z ]{2,30}$/
    const Employees = JSON.parse(window.sessionStorage.getItem('Employees'));
    
    if(form.code.value=="" || !codeRegEx.test(form.code.value)){
        alert("Invalid Code [Only alphanumeric allowed]");
        return false;
    }else if(form.name.value=="" || !nameRegex.test(form.name.value)){
        alert("Invalid Name format");
        return false;
    }else if(form.email.value=="" || !emailRegEx.test(form.email.value)){
        alert("Invalid Email format");
        return false;
    }else if(form.des.value=="" || !desRegEx.test(form.des.value)){
        alert("Invalid Designation Format");
        return false;
    }


    for(let emp of Employees){
        if(emp['code']==form.code.value){
            alert("User with given code exists");
            return false;
        }
    }
    return true;
    
}


//function to process form on adding user
const processForm = () =>{
    document.getElementById('Add-employee').addEventListener('submit',(e)=>{
        e.preventDefault();
    });

    const form = document.getElementById('Add-employee');

    
    const Employees = JSON.parse(window.sessionStorage.getItem('Employees'));
    console.log(Employees);

    if(!addFormValidation(form)){
        return;
    }else{
        const new_employee = {

            name : form.name!=undefined && form.name.value,
            code : form.code!=undefined && form.code.value,
            email : form.email!=undefined && form.email.value,
            gender : form.gender!=undefined && form.gender.value,
            designation : form.des!=undefined && form.des.value,
            dob : form.dob!=undefined && form.dob.value,
        };
    
        Employees.push(new_employee);
        window.sessionStorage.setItem('Employees',JSON.stringify(Employees));
        console.log(window.sessionStorage);
    
        window.location.pathname="../index.html";
    
    }
}

//Validation function for update user
const updateFormValidation = (form)=>{
    let codeRegEx = /^[a-zA-Z0-9]*$/;
    let emailRegEx = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    let nameRegex = /^[a-zA-Z ]*$/;
    let desRegEx = /^[a-zA-Z ]{2,30}$/
    
    if(form.code.value!="" && !codeRegEx.test(form.code.value)){
        alert("Invalid Code [Only alphanumeric allowed]");
        return false;
    }else if(form.name.value!="" && !nameRegex.test(form.name.value)){
        alert("Invalid Name format");
        return false;
    }else if(form.email.value!="" && !emailRegEx.test(form.email.value)){
        alert("Invalid Email format");
        return false;
    }else if(form.des.value!="" && !desRegEx.test(form.des.value)){
        alert("Invalid Designation Format");
        return false;
    }

    return true;
    
}

//Function to process update user form
const processUpdateForm = () =>{
    document.getElementById('form').addEventListener('submit',(e)=>{
        e.preventDefault();
    });

    const form = document.getElementById("form");
    const Employees = JSON.parse(window.sessionStorage.getItem('Employees'));
    console.log(typeof Employees);
    let isValid = false;

    for(let employee of Employees){
        if(employee['code']==form.code.value){
            isValid = true;
            if(!updateFormValidation(form)){
                console.log("Not validated");
                return;                
            }else{
                employee['name'] = (form.name.value!="")?form.name.value:employee['name'];
                employee['code'] = employee['code'];
                employee['email'] = (form.email.value!="")?form.email.value:employee['email'];
                employee['gender'] = (form.gender.value!="")?form.gender.value:employee['gender'];
                employee['designation'] = (form.des.value!="")?form.des.value:employee['designation'];
                employee['dob'] = (form.dob.value!="")?form.dob.value:employee['dob'];             
            }
            break;
        }
    }

    if(!isValid){
        alert("Employee Doesn't Exist");
    }else{
        window.sessionStorage.setItem('Employees',JSON.stringify(Employees));
        console.log(window.sessionStorage);

        if(window.sessionStorage.getItem("update_code")!=undefined){
            window.sessionStorage.removeItem('update_code');
        }
    
        window.location.pathname="../index.html";
    }

}
