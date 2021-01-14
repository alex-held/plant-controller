import {Component, OnInit} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {GrowConfig} from './growconfig.model';
import {MessageService} from "primeng/api";
import {GrowConfigService} from "./growconfig.service";

@Component({
    selector: 'app-growconfig',
    templateUrl: './growconfig.component.html',
    styleUrls: ['./growconfig.component.scss']
})
export class GrowConfigComponent implements OnInit {

    growconfigs: GrowConfig[];
    growconfig: GrowConfig;
    growconfigDialog: boolean;
    submitted: boolean;

    constructor(private growConfigService: GrowConfigService, private httpClient: HttpClient, private messageService: MessageService) {
    }


    ngOnInit(): void {
        this.growconfig = {};
        this.loadData();
    }

    loadData() {
        console.log('loading data...');
        this.growConfigService.read().subscribe(data => this.growconfigs = data);
    }

    delete(id: number) {
        this.growConfigService.delete(id).subscribe(() => console.log('deleted'))
        this.loadData();
    }

    edit(growconfig: GrowConfig) {
        this.growconfig = {...growconfig};
        this.growconfigDialog = true;
    }

    save() {
        this.submitted = true;

        if (this.growconfig.Id === null || this.growconfig.Id === undefined) {
            console.log(JSON.stringify(this.growconfig))
            this.growConfigService.create(this.growconfig)
                .subscribe(value => console.log(JSON.stringify(value)));
            this.messageService.add({
                severity: 'success',
                summary: 'Successful',
                detail: 'GrowConfig created',
                life: 3000
            });
        } else {

            console.log(JSON.stringify(this.growconfig))
            this.growConfigService.update(this.growconfig)
                .subscribe(value => console.log(JSON.stringify(value)));
            this.messageService.add({
                severity: 'success',
                summary: 'Successful',
                detail: 'GrowConfig updated',
                life: 3000
            });

        }

        this.loadData();
        this.growconfigDialog = false;
        this.growconfig = {};
    }

    openNew() {
        this.growconfig = {};
        this.submitted = false;
        this.growconfigDialog = true;
    }

    hideDialog() {
        this.growconfigDialog = false;
        this.submitted = false;
    }

}
