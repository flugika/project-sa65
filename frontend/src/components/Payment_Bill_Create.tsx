import React, { useState, useEffect } from 'react';
import Button from '@mui/material/Button';
import Container from '@mui/material/Container';
import Box from '@mui/material/Box';
import Paper from '@mui/material/Paper';
import Grid from '@mui/material/Grid';
import { styled } from '@mui/material/styles';
import TextField from '@mui/material/TextField';
import InputAdornment from '@mui/material/InputAdornment';
import Clock from 'react-live-clock';
import FormControl from '@mui/material/FormControl';
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { Link as RouterLink } from "react-router-dom";
import Snackbar from "@mui/material/Snackbar";
import Select, { SelectChangeEvent } from "@mui/material/Select";

import { BookingInterface } from "../interfaces/IBooking";
import { PaymentBillInterface } from "../interfaces/IPaymentBill";
import { RoomPriceInterface } from "../interfaces/IRoomPrice";
import { EmployeeInterface } from "../interfaces/IEmployee";
import { RoomInterface } from "../interfaces/IRoom";
import { StudentInterface } from '../interfaces/IStudent';
import { SemesterInterface } from '../interfaces/ISemester';

import {
  GetEmployee,
  GetRoom,
  GetRoom_Price,
  GetBooking,
  GetStudent,
  GetSemester,
  Payment_Bills,
} from "../services/HttpClientService";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

export default function Payment_Bill_Create() {

  const [room_Price, setRoom_Price] = useState<RoomPriceInterface[]>([]);
  const [student, setStudent] = useState<StudentInterface>();
  const [employee, setEmployee] = useState<EmployeeInterface[]>([]);
  const [booking, setBooking] = useState<BookingInterface[]>([]);
  const [room, setRoom] = useState<RoomInterface[]>([]);
  const [semester, setSemester] = useState<SemesterInterface[]>([]);
  const [payment_bill, setPayment_Bill] = useState<PaymentBillInterface>({});
  const [Electric_Bill, setElectric_Bill] = useState<string>()
  const [Water_Bill, setWater_Bill] = useState<string>()

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof payment_bill;
    setPayment_Bill({
      ...payment_bill,
      [name]: event.target.value,
    });
  };

  /* function parseJwt(token: string) {
    if (!token) { return; }
    const base64Url = token.split('.')[1];
    const base64 = base64Url.replace('-', '+').replace('_', '/');
    return JSON.parse(window.atob(base64));
  } */
  
  /* const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbXBsb3llZV9FbWFpbCI6IkBtaW4uY29tIiwiZXhwIjoxNjY2NzI1ODI3LCJpc3MiOiJBdXRoU2VydmljZSJ9.mHzhuymifQMEAPe9xGNGai958syAHm-si41G8qSpLtY";
  const current_employee = parseJwt(token);

  console.log(current_employee.Employee_Email); */

  const getRoom_Price = async () => {
    let res = await GetRoom_Price();
    if (res) {
      setRoom_Price(res);
    }
  };

  const getEmployee = async () => {
    let res = await GetEmployee();
    if (res) {
      setEmployee(res);
    }
  };

  const getRoom = async () => {
    let res = await GetRoom();
    if (res) {
      setRoom(res);
    }
  };

  const getStudent = async () => {
    let res = await GetStudent();
    if (res) {
      setStudent(res);
    }
  };

  const getBooking = async () => {
    let res = await GetBooking();
    if (res) {
      setBooking(res);
    }
  };

  const getSemester = async () => {
    let res = await GetSemester();
    if (res) {
      setSemester(res);
    }
  };

  useEffect(() => {
    getRoom_Price();
    getEmployee();
    getRoom();
    getStudent();
    getBooking();
    getSemester();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  const Cal_PaymentBalance = (Electric_Bill: number, Water_Bill: number) => {
    return Electric_Bill + Water_Bill + 4500;
  }

  async function generate() {
    let data = {
      BookingID: convertType(payment_bill.BookingID),
      EmployeeID: convertType(payment_bill.EmployeeID),
      SemesterID: convertType(payment_bill.SemesterID),
      Billing_Date: new Date(), // สร้างค่าใหม่เพื่อให้เวลาเป็นปัจจุบันตอนบันทึก
      Electric_Bill: typeof Electric_Bill == "string" ? parseInt(Electric_Bill) : 0,
      Water_Bill: typeof Water_Bill == "string" ? parseInt(Water_Bill) : 0,
      Payment_Balance: Cal_PaymentBalance(Number(Electric_Bill), Number(Water_Bill)),
    };

    let res = await Payment_Bills(data);
    if (res) {
      setSuccess(true);
    } else {
      setError(true);
    }
  }

  const Item = styled(Paper)(({ theme }) => ({
    backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
    ...theme.typography.body2,
    padding: theme.spacing(1),
    textAlign: 'center',
    color: theme.palette.text.secondary,
  }));

  return (
    <div>
      <Snackbar
        open={success}
        autoHideDuration={3000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar
        open={error}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>

      {/* CONTAINER =============================== */}
      <Container maxWidth="md">
        <Paper >
          <Box
            display={"flex"}
            sx={{
              marginTop: 2,
              paddingX: 2,
              paddingY: 1,
            }}>
            <Grid container spacing={2}>
              <Grid item xs={6}>
                <Box bgcolor="success.main" style={{border: '1px solid black',borderRadius: '50px', marginTop: 20,}}>
                  <h2 style={{color: "black", textAlign: "center"}} >สร้างใบเสร็จชำระค่าใช้จ่าย</h2>
                  </Box>
              </Grid>
              <Grid item xs={6} >
                <Item style={{
                  marginTop: 75,
                  marginLeft: 100,
                }}
                >
                  <p style={{
                    height: 3,
                    textAlign: "center",
                    paddingBottom: 6,
                    marginTop: 6,
                    paddingRight: 6,
                  }}
                  >
                    <Clock format={'ddd, YYYY-MM-DD HH:mm:ss'} ticking={true} timezone={'Asia/bangkok'} /></p>
                </Item>
              </Grid>
            </Grid>
          </Box>
          <hr />
          <Grid container spacing={2}>

            <Grid item xs={12}>
            </Grid>
            <Grid item xs={1.5}>
            </Grid>
            <Grid item xs={8}>
              <Grid container spacing={2} style={{paddingRight: 20, paddingBottom: 20}}>
              <Grid item xs={4}>
                  <p style={{
                    textAlign: "right",
                    fontWeight: "bold",
                    }}
                    >พนักงาน
                  </p>
                </Grid>
                <Grid item xs={8} >
                  <FormControl fullWidth variant="outlined">
                    <Select
                      native
                      value={payment_bill.EmployeeID + ""}
                      onChange={handleChange}
                      inputProps={{
                        name: "EmployeeID",
                      }}
                    >
                      <option aria-label="None" value="">
                        กรุณาเลือกพนักงาน
                      </option>
                      {employee.map((item: EmployeeInterface) => (
                        <option value={item.ID} key={item.ID}>
                          {item.Employee_Name}
                        </option>
                      ))}
                    </Select>
                  </FormControl>
                </Grid>
                
                <Grid item xs={4}>
                  <p style={{
                    textAlign: "right",
                    fontWeight: "bold",
                    }}
                    >รหัสการจอง
                  </p>
                </Grid>
                <Grid item xs={8} >
                  <FormControl fullWidth variant="outlined">
                    <Select
                      native
                      value={payment_bill.BookingID + ""}
                      onChange={handleChange}
                      inputProps={{
                        name: "BookingID",
                      }}
                    >
                      <option aria-label="None" value="">
                        กรุณาเลือกรหัสการจอง
                      </option>
                      {booking.map((item: BookingInterface) => (
                        <option value={item.ID} key={item.ID}>
                          {item.ID}
                        </option>
                      ))}
                    </Select>
                  </FormControl>
                </Grid>

                <Grid item xs={4}>
                  <p style={{
                    textAlign: "right",
                    fontWeight: "bold",
                    }}
                    >ภาคการศึกษา
                  </p>
                </Grid>
                <Grid item xs={8} >
                  <FormControl fullWidth variant="outlined">
                    <Select
                      native
                      value={payment_bill.SemesterID + ""}
                      onChange={handleChange}
                      inputProps={{
                        name: "SemesterID",
                      }}
                    >
                      <option aria-label="None" value="">
                        กรุณาเลือกภาคการศึกษา
                      </option>
                      {semester.map((item: SemesterInterface) => (
                        <option value={item.ID} key={item.ID}>
                          {item.Semester}
                        </option>
                      ))}
                    </Select>
                  </FormControl>
                </Grid>

                <Grid item xs={4}>
                  <p style={{
                    textAlign: "right",
                    fontWeight: "bold",
                    }}
                    >ค่าไฟ
                  </p>
                </Grid>
                <Grid item xs={8}>
                  <TextField 
                    fullWidth
                    id="Electric_Bill"
                    defaultValue=""
                    type="number"
                    variant="outlined"
                    size="small"
                    InputProps={{
                      inputProps: {min: 0},
                      endAdornment: <InputAdornment position="end"><h5>บาท</h5></InputAdornment>
                    }}
                    autoComplete="off"
                    onChange={(event) => setElectric_Bill(event.target.value)}
                  />
                </Grid>

                <Grid item xs={4}>
                  <p style={{
                    textAlign: "right",
                    fontWeight: "bold",
                    }}
                    >ค่าน้ำ
                  </p>
                </Grid>
                <Grid item xs={8}>
                  <TextField 
                    fullWidth
                    id="Warter_Bill"
                    defaultValue=""
                    type="number"
                    variant="outlined"
                    size="small"
                    InputProps={{
                      inputProps: {min: 0},
                      endAdornment: <InputAdornment position="end"><h5>บาท</h5></InputAdornment>
                    }}
                    autoComplete="off"
                    onChange={(event) => setWater_Bill(event.target.value)}
                  />
                </Grid>
                
                <Grid item xs={4}></Grid>
                <Grid item xs={4}>
                  <Button
                    component={RouterLink}
                    to="/payment_bills"
                    variant="contained"
                    color="inherit"
                  >
                    กลับ
                  </Button>
                </Grid>
                <Grid item xs={4}>
                  <Button 
                    style={{fontWeight: "bold", fontSize:"sm"}}
                    fullWidth
                    variant="contained"
                    onClick={generate}
                    >สร้าง
                  </Button>
                </Grid>
              </Grid>
            </Grid>
          </Grid>
        </Paper>
      </Container>
    </div>
  );
}