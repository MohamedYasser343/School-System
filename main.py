from flask import Flask, render_template, request, send_from_directory
import os
import pandas as pd
from fpdf import FPDF
import mysql.connector

filename = "FPdf"

def getReport():
    mydb = mysql.connector.connect(
    host="192.168.103.143",
    user="snap",
    password="snap",
    database="snap",
    )

    db = mydb.cursor()

    db.execute("SELECT * FROM `StudentViolations` WHERE ID = 2520060;")
    result = db.fetchone()
    stdName = str(result[0])
    print(stdName)

    Height = 297
    Width = 210
    data_frame = pd.read_csv("static/students.csv")
    filename = "FPdf"
    religons = []
    for x in data_frame['stdreligon']:
        if x == 'Christian':
            religons.append(0)
        else:
            religons.append(1)
    data_frame['stdreligon'] = religons
    export = data_frame['stdreligon'].plot()
    export.get_figure().savefig(f"./temp/{filename}.png", dpi=300)
    pdf = FPDF()

    pdf.add_page()
    pdf.set_font('Arial', 'B', 16)
    pdf.set_y(-30)
    pdf.cell(w = 0, h = 0, txt = f"Report Of Student: {stdName}", align = "center")
    pdf.set_y(0)
    pdf.image("./reportImages/Logo.png", 88, 5, 40, 40)
    pdf.image(f"./temp/{filename}.png", 10, 60, Width-10)
    pdf.output(f'./static/{filename}.pdf')

app = Flask(__name__)

@app.route('/GetReport')

def index():
    return render_template('GeneratedReport.html')

@app.route("/P/")

def pdf():
    filepath = os.path.abspath(os.getcwd()) + '/static'
    return send_from_directory(filepath, f'{filename}.pdf')
