import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { PaymentBillInterface } from "../interfaces/IPaymentBill";
import { GetPayment_Bills } from "../services/HttpClientService";

function Payment_Bill() {
  const [payment_bills, setPayment_Bills] = useState<PaymentBillInterface[]>([]);

  useEffect(() => {
    getPayment_Bills();
  }, []);

  const getPayment_Bills = async () => {
    let res = await GetPayment_Bills();
    if (res) {
      setPayment_Bills(res);
    } 
  };

  const columns: GridColDef[] = [
    { field: "ID", headerName: "ลำดับ", width: 50 },
    { field: "BookingID", headerName: "รหัสการจอง", width: 100 },
    { field: "SemesterID", headerName: "รหัสภาคการศึกษา", width: 150 },
    {
      field: "Electric_Bill",
      headerName: "ค่าไฟ",
      width: 100,
      valueFormatter: (params) => params.value.Name,
    },
    {
      field: "Water_Bill",
      headerName: "ค่าน้ำ",
      width: 100,
      valueFormatter: (params) => params.value.Value,
    },
    {
      field: "Payment_Balance",
      headerName: "รวมค่าใช้จ่าย",
      width: 150,
      valueFormatter: (params) => params.value.Value,
    },
    { field: "Billing_Date", headerName: "วันที่และเวลา", width: 300 },
  ];

  return (
    <div>
      <Container maxWidth="md">
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ข้อมูลใบเสร็จค่าใช้จ่าย
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/payment_bills/create"
              variant="contained"
              color="primary"
            >
              สร้างใบเสร็จ
            </Button>
          </Box>
        </Box>
        <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
          <DataGrid
            rows={payment_bills}
            getRowId={(row) => row.ID}
            columns={columns}
            pageSize={5}
            rowsPerPageOptions={[5]}
          />
        </div>
      </Container>
    </div>
  );
}

export default Payment_Bill;