import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { environment } from '../environments/environment';
import { ChangeDetectorRef } from '@angular/core';
import { MonacoEditorModule, NGX_MONACO_EDITOR_CONFIG } from 'ngx-monaco-editor-v2';
import { CommonModule } from '@angular/common';


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [FormsModule, MonacoEditorModule, CommonModule],
  templateUrl: './app.html',
  providers: [
  {
    provide: NGX_MONACO_EDITOR_CONFIG,
    useValue: {
      baseUrl: 'assets/monaco/vs',
    },
  },
],
})
export class AppComponent implements OnInit {
  isLoading = false;
  executionTime: string = '';  
  status: string = '';         
  templates: any = {
    python: "print('Hello World')",

    java: `public class Main {
            public static void main(String[] args) {
              System.out.println("Hello World");
            }
          }`,

    cpp: `#include <iostream>
          using namespace std;

          int main() {
            cout << "Hello World";
            return 0;
          }`
  };

  code = "print('Hello from Angular')";
  language = 'python';
  output = '';
  editorOptions = {
  theme: 'vs-dark',
  language: 'python',
  automaticLayout: true
  };
  onLanguageChange() {
  this.code = this.templates[this.language];

  this.editorOptions = {
    ...this.editorOptions,
    language: this.language === 'cpp' ? 'cpp' : this.language
  };
}

  API_URL = `${environment.apiUrl}/run/`;

  constructor(private http: HttpClient, private cd: ChangeDetectorRef) {}

  ngOnInit() {
    // Warmup request
    this.http.post(this.API_URL, {
      code: "print('warmup')",
      language: "python"
    }).subscribe();
  }


  runCode() {
  this.isLoading = true;
  this.output = "";

  this.http.post<any>(this.API_URL, {
    code: this.code,
    language: this.language
  }).subscribe({
    next: (res) => {
      console.log('API Response:', res);
      this.output = res.output;
      // this.output = "TEST OUTPUT"
      this.executionTime = res.time;
      this.status = res.status;
      this.isLoading = false;

      this.cd.detectChanges();
    },
    error: () => {
      this.output = "Error occurred";
      this.isLoading = false;
    }
  });
}
}