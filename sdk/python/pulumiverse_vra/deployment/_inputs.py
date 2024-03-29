# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'DeploymentExpenseArgs',
    'DeploymentLastRequestArgs',
    'DeploymentProjectArgs',
    'DeploymentResourceArgs',
    'DeploymentResourceExpenseArgs',
]

@pulumi.input_type
class DeploymentExpenseArgs:
    def __init__(__self__, *,
                 additional_expense: Optional[pulumi.Input[float]] = None,
                 code: Optional[pulumi.Input[str]] = None,
                 compute_expense: Optional[pulumi.Input[float]] = None,
                 last_update_time: Optional[pulumi.Input[str]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 network_expense: Optional[pulumi.Input[float]] = None,
                 storage_expense: Optional[pulumi.Input[float]] = None,
                 total_expense: Optional[pulumi.Input[float]] = None,
                 unit: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[float] additional_expense: Additional expense incurred for the resource.
        :param pulumi.Input[str] code: Expense sync message code if any.
        :param pulumi.Input[float] compute_expense: Compute expense of the entity.
        :param pulumi.Input[str] last_update_time: Last expense sync time.
        :param pulumi.Input[str] message: Expense sync message if any.
        :param pulumi.Input[float] network_expense: Network expense of the entity.
        :param pulumi.Input[float] storage_expense: Storage expense of the entity.
        :param pulumi.Input[float] total_expense: Total expense of the entity.
        :param pulumi.Input[str] unit: Monetary unit.
        """
        if additional_expense is not None:
            pulumi.set(__self__, "additional_expense", additional_expense)
        if code is not None:
            pulumi.set(__self__, "code", code)
        if compute_expense is not None:
            pulumi.set(__self__, "compute_expense", compute_expense)
        if last_update_time is not None:
            pulumi.set(__self__, "last_update_time", last_update_time)
        if message is not None:
            pulumi.set(__self__, "message", message)
        if network_expense is not None:
            pulumi.set(__self__, "network_expense", network_expense)
        if storage_expense is not None:
            pulumi.set(__self__, "storage_expense", storage_expense)
        if total_expense is not None:
            pulumi.set(__self__, "total_expense", total_expense)
        if unit is not None:
            pulumi.set(__self__, "unit", unit)

    @property
    @pulumi.getter(name="additionalExpense")
    def additional_expense(self) -> Optional[pulumi.Input[float]]:
        """
        Additional expense incurred for the resource.
        """
        return pulumi.get(self, "additional_expense")

    @additional_expense.setter
    def additional_expense(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "additional_expense", value)

    @property
    @pulumi.getter
    def code(self) -> Optional[pulumi.Input[str]]:
        """
        Expense sync message code if any.
        """
        return pulumi.get(self, "code")

    @code.setter
    def code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "code", value)

    @property
    @pulumi.getter(name="computeExpense")
    def compute_expense(self) -> Optional[pulumi.Input[float]]:
        """
        Compute expense of the entity.
        """
        return pulumi.get(self, "compute_expense")

    @compute_expense.setter
    def compute_expense(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "compute_expense", value)

    @property
    @pulumi.getter(name="lastUpdateTime")
    def last_update_time(self) -> Optional[pulumi.Input[str]]:
        """
        Last expense sync time.
        """
        return pulumi.get(self, "last_update_time")

    @last_update_time.setter
    def last_update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_update_time", value)

    @property
    @pulumi.getter
    def message(self) -> Optional[pulumi.Input[str]]:
        """
        Expense sync message if any.
        """
        return pulumi.get(self, "message")

    @message.setter
    def message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message", value)

    @property
    @pulumi.getter(name="networkExpense")
    def network_expense(self) -> Optional[pulumi.Input[float]]:
        """
        Network expense of the entity.
        """
        return pulumi.get(self, "network_expense")

    @network_expense.setter
    def network_expense(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "network_expense", value)

    @property
    @pulumi.getter(name="storageExpense")
    def storage_expense(self) -> Optional[pulumi.Input[float]]:
        """
        Storage expense of the entity.
        """
        return pulumi.get(self, "storage_expense")

    @storage_expense.setter
    def storage_expense(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "storage_expense", value)

    @property
    @pulumi.getter(name="totalExpense")
    def total_expense(self) -> Optional[pulumi.Input[float]]:
        """
        Total expense of the entity.
        """
        return pulumi.get(self, "total_expense")

    @total_expense.setter
    def total_expense(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "total_expense", value)

    @property
    @pulumi.getter
    def unit(self) -> Optional[pulumi.Input[str]]:
        """
        Monetary unit.
        """
        return pulumi.get(self, "unit")

    @unit.setter
    def unit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unit", value)


@pulumi.input_type
class DeploymentLastRequestArgs:
    def __init__(__self__, *,
                 action_id: Optional[pulumi.Input[str]] = None,
                 approved_at: Optional[pulumi.Input[str]] = None,
                 blueprint_id: Optional[pulumi.Input[str]] = None,
                 cancelable: Optional[pulumi.Input[bool]] = None,
                 catalog_item_id: Optional[pulumi.Input[str]] = None,
                 completed_at: Optional[pulumi.Input[str]] = None,
                 completed_tasks: Optional[pulumi.Input[int]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 details: Optional[pulumi.Input[str]] = None,
                 dismissed: Optional[pulumi.Input[bool]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 initialized_at: Optional[pulumi.Input[str]] = None,
                 inputs: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 outputs: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 requested_by: Optional[pulumi.Input[str]] = None,
                 resource_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 total_tasks: Optional[pulumi.Input[int]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] action_id: Identifier of the requested action.
        :param pulumi.Input[str] approved_at: Time at which the request was approved.
        :param pulumi.Input[str] blueprint_id: Identifier of the requested blueprint in the form ‘UUID:version’.
        :param pulumi.Input[bool] cancelable: Indicates whether request can be canceled or not.
        :param pulumi.Input[str] catalog_item_id: The id of the vRA catalog item to request the deployment. Conflicts with `blueprint_id` and `blueprint_content`.
        :param pulumi.Input[str] completed_at: Time at which the request completed.
        :param pulumi.Input[int] completed_tasks: The number of tasks completed while fulfilling this request.
        :param pulumi.Input[str] created_at: Creation time (e.g. date format ‘2019-07-13T23:16:49.310Z’).
        :param pulumi.Input[str] details: Longer user-friendly details of the request.
        :param pulumi.Input[bool] dismissed: Indicates whether request is in dismissed state.
        :param pulumi.Input[str] id: Unique identifier of the resource.
        :param pulumi.Input[str] initialized_at: Time at which the request was initialized.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] inputs: Inputs provided by the user. For inputs including those with default values, refer to `inputs_including_defaults`.
        :param pulumi.Input[str] name: A human-friendly name used as an identifier in APIs that support this option.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] outputs: Request outputs.
        :param pulumi.Input[str] requested_by: The user that initiated the request.
        :param pulumi.Input[str] status: Deployment status. Supported values are: `CREATE_SUCCESSFUL`, `CREATE_INPROGRESS`, `CREATE_FAILED`, `UPDATE_SUCCESSFUL`, `UPDATE_INPROGRESS`, `UPDATE_FAILED`, `DELETE_SUCCESSFUL`, `DELETE_INPROGRESS`, `DELETE_FAILED`, `ACTION_SUCCESSFUL`, `ACTION_INPROGRESS`, `ACTION_FAILED`.
        :param pulumi.Input[str] updated_at: Last update time (e.g. date format ‘2019-07-13T23:16:49.310Z’).
        """
        if action_id is not None:
            pulumi.set(__self__, "action_id", action_id)
        if approved_at is not None:
            pulumi.set(__self__, "approved_at", approved_at)
        if blueprint_id is not None:
            pulumi.set(__self__, "blueprint_id", blueprint_id)
        if cancelable is not None:
            pulumi.set(__self__, "cancelable", cancelable)
        if catalog_item_id is not None:
            pulumi.set(__self__, "catalog_item_id", catalog_item_id)
        if completed_at is not None:
            pulumi.set(__self__, "completed_at", completed_at)
        if completed_tasks is not None:
            pulumi.set(__self__, "completed_tasks", completed_tasks)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if details is not None:
            pulumi.set(__self__, "details", details)
        if dismissed is not None:
            pulumi.set(__self__, "dismissed", dismissed)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if initialized_at is not None:
            pulumi.set(__self__, "initialized_at", initialized_at)
        if inputs is not None:
            pulumi.set(__self__, "inputs", inputs)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if outputs is not None:
            pulumi.set(__self__, "outputs", outputs)
        if requested_by is not None:
            pulumi.set(__self__, "requested_by", requested_by)
        if resource_ids is not None:
            pulumi.set(__self__, "resource_ids", resource_ids)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if total_tasks is not None:
            pulumi.set(__self__, "total_tasks", total_tasks)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="actionId")
    def action_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the requested action.
        """
        return pulumi.get(self, "action_id")

    @action_id.setter
    def action_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action_id", value)

    @property
    @pulumi.getter(name="approvedAt")
    def approved_at(self) -> Optional[pulumi.Input[str]]:
        """
        Time at which the request was approved.
        """
        return pulumi.get(self, "approved_at")

    @approved_at.setter
    def approved_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "approved_at", value)

    @property
    @pulumi.getter(name="blueprintId")
    def blueprint_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the requested blueprint in the form ‘UUID:version’.
        """
        return pulumi.get(self, "blueprint_id")

    @blueprint_id.setter
    def blueprint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "blueprint_id", value)

    @property
    @pulumi.getter
    def cancelable(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether request can be canceled or not.
        """
        return pulumi.get(self, "cancelable")

    @cancelable.setter
    def cancelable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "cancelable", value)

    @property
    @pulumi.getter(name="catalogItemId")
    def catalog_item_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the vRA catalog item to request the deployment. Conflicts with `blueprint_id` and `blueprint_content`.
        """
        return pulumi.get(self, "catalog_item_id")

    @catalog_item_id.setter
    def catalog_item_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "catalog_item_id", value)

    @property
    @pulumi.getter(name="completedAt")
    def completed_at(self) -> Optional[pulumi.Input[str]]:
        """
        Time at which the request completed.
        """
        return pulumi.get(self, "completed_at")

    @completed_at.setter
    def completed_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "completed_at", value)

    @property
    @pulumi.getter(name="completedTasks")
    def completed_tasks(self) -> Optional[pulumi.Input[int]]:
        """
        The number of tasks completed while fulfilling this request.
        """
        return pulumi.get(self, "completed_tasks")

    @completed_tasks.setter
    def completed_tasks(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "completed_tasks", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Creation time (e.g. date format ‘2019-07-13T23:16:49.310Z’).
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def details(self) -> Optional[pulumi.Input[str]]:
        """
        Longer user-friendly details of the request.
        """
        return pulumi.get(self, "details")

    @details.setter
    def details(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "details", value)

    @property
    @pulumi.getter
    def dismissed(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether request is in dismissed state.
        """
        return pulumi.get(self, "dismissed")

    @dismissed.setter
    def dismissed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dismissed", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier of the resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="initializedAt")
    def initialized_at(self) -> Optional[pulumi.Input[str]]:
        """
        Time at which the request was initialized.
        """
        return pulumi.get(self, "initialized_at")

    @initialized_at.setter
    def initialized_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "initialized_at", value)

    @property
    @pulumi.getter
    def inputs(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Inputs provided by the user. For inputs including those with default values, refer to `inputs_including_defaults`.
        """
        return pulumi.get(self, "inputs")

    @inputs.setter
    def inputs(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "inputs", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def outputs(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Request outputs.
        """
        return pulumi.get(self, "outputs")

    @outputs.setter
    def outputs(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "outputs", value)

    @property
    @pulumi.getter(name="requestedBy")
    def requested_by(self) -> Optional[pulumi.Input[str]]:
        """
        The user that initiated the request.
        """
        return pulumi.get(self, "requested_by")

    @requested_by.setter
    def requested_by(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "requested_by", value)

    @property
    @pulumi.getter(name="resourceIds")
    def resource_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "resource_ids")

    @resource_ids.setter
    def resource_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resource_ids", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Deployment status. Supported values are: `CREATE_SUCCESSFUL`, `CREATE_INPROGRESS`, `CREATE_FAILED`, `UPDATE_SUCCESSFUL`, `UPDATE_INPROGRESS`, `UPDATE_FAILED`, `DELETE_SUCCESSFUL`, `DELETE_INPROGRESS`, `DELETE_FAILED`, `ACTION_SUCCESSFUL`, `ACTION_INPROGRESS`, `ACTION_FAILED`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="totalTasks")
    def total_tasks(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "total_tasks")

    @total_tasks.setter
    def total_tasks(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "total_tasks", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        Last update time (e.g. date format ‘2019-07-13T23:16:49.310Z’).
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


@pulumi.input_type
class DeploymentProjectArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] description: A human-friendly description.
        :param pulumi.Input[str] id: Unique identifier of the resource.
        :param pulumi.Input[str] name: A human-friendly name used as an identifier in APIs that support this option.
        :param pulumi.Input[str] version: Version of the entity, if applicable.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier of the resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        Version of the entity, if applicable.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class DeploymentResourceArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 name: pulumi.Input[str],
                 created_at: Optional[pulumi.Input[str]] = None,
                 depends_ons: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 expenses: Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentResourceExpenseArgs']]]] = None,
                 properties_json: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 sync_status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: Unique identifier of the resource.
        :param pulumi.Input[str] name: A human-friendly name used as an identifier in APIs that support this option.
        :param pulumi.Input[str] created_at: Creation time (e.g. date format ‘2019-07-13T23:16:49.310Z’).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] depends_ons: A list of other resources this resource depends on.
        :param pulumi.Input[str] description: A human-friendly description.
        :param pulumi.Input[Sequence[pulumi.Input['DeploymentResourceExpenseArgs']]] expenses: Expense incurred for the deployment.
        :param pulumi.Input[str] properties_json: List of properties in the encoded JSON string format.
        :param pulumi.Input[str] state: The current state of the resource. Supported values are `PARTIAL`, `TAINTED`, `OK.`
        :param pulumi.Input[str] sync_status: The current sync status. Supported values are `SUCCESS`, `MISSING`, `STALE`.
        :param pulumi.Input[str] type: Type of the resource.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if depends_ons is not None:
            pulumi.set(__self__, "depends_ons", depends_ons)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if expenses is not None:
            pulumi.set(__self__, "expenses", expenses)
        if properties_json is not None:
            pulumi.set(__self__, "properties_json", properties_json)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if sync_status is not None:
            pulumi.set(__self__, "sync_status", sync_status)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        Unique identifier of the resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        A human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Creation time (e.g. date format ‘2019-07-13T23:16:49.310Z’).
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="dependsOns")
    def depends_ons(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of other resources this resource depends on.
        """
        return pulumi.get(self, "depends_ons")

    @depends_ons.setter
    def depends_ons(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "depends_ons", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def expenses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentResourceExpenseArgs']]]]:
        """
        Expense incurred for the deployment.
        """
        return pulumi.get(self, "expenses")

    @expenses.setter
    def expenses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentResourceExpenseArgs']]]]):
        pulumi.set(self, "expenses", value)

    @property
    @pulumi.getter(name="propertiesJson")
    def properties_json(self) -> Optional[pulumi.Input[str]]:
        """
        List of properties in the encoded JSON string format.
        """
        return pulumi.get(self, "properties_json")

    @properties_json.setter
    def properties_json(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "properties_json", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The current state of the resource. Supported values are `PARTIAL`, `TAINTED`, `OK.`
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="syncStatus")
    def sync_status(self) -> Optional[pulumi.Input[str]]:
        """
        The current sync status. Supported values are `SUCCESS`, `MISSING`, `STALE`.
        """
        return pulumi.get(self, "sync_status")

    @sync_status.setter
    def sync_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sync_status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the resource.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class DeploymentResourceExpenseArgs:
    def __init__(__self__, *,
                 additional_expense: Optional[pulumi.Input[float]] = None,
                 code: Optional[pulumi.Input[str]] = None,
                 compute_expense: Optional[pulumi.Input[float]] = None,
                 last_update_time: Optional[pulumi.Input[str]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 network_expense: Optional[pulumi.Input[float]] = None,
                 storage_expense: Optional[pulumi.Input[float]] = None,
                 total_expense: Optional[pulumi.Input[float]] = None,
                 unit: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[float] additional_expense: Additional expense incurred for the resource.
        :param pulumi.Input[str] code: Expense sync message code if any.
        :param pulumi.Input[float] compute_expense: Compute expense of the entity.
        :param pulumi.Input[str] last_update_time: Last expense sync time.
        :param pulumi.Input[str] message: Expense sync message if any.
        :param pulumi.Input[float] network_expense: Network expense of the entity.
        :param pulumi.Input[float] storage_expense: Storage expense of the entity.
        :param pulumi.Input[float] total_expense: Total expense of the entity.
        :param pulumi.Input[str] unit: Monetary unit.
        """
        if additional_expense is not None:
            pulumi.set(__self__, "additional_expense", additional_expense)
        if code is not None:
            pulumi.set(__self__, "code", code)
        if compute_expense is not None:
            pulumi.set(__self__, "compute_expense", compute_expense)
        if last_update_time is not None:
            pulumi.set(__self__, "last_update_time", last_update_time)
        if message is not None:
            pulumi.set(__self__, "message", message)
        if network_expense is not None:
            pulumi.set(__self__, "network_expense", network_expense)
        if storage_expense is not None:
            pulumi.set(__self__, "storage_expense", storage_expense)
        if total_expense is not None:
            pulumi.set(__self__, "total_expense", total_expense)
        if unit is not None:
            pulumi.set(__self__, "unit", unit)

    @property
    @pulumi.getter(name="additionalExpense")
    def additional_expense(self) -> Optional[pulumi.Input[float]]:
        """
        Additional expense incurred for the resource.
        """
        return pulumi.get(self, "additional_expense")

    @additional_expense.setter
    def additional_expense(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "additional_expense", value)

    @property
    @pulumi.getter
    def code(self) -> Optional[pulumi.Input[str]]:
        """
        Expense sync message code if any.
        """
        return pulumi.get(self, "code")

    @code.setter
    def code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "code", value)

    @property
    @pulumi.getter(name="computeExpense")
    def compute_expense(self) -> Optional[pulumi.Input[float]]:
        """
        Compute expense of the entity.
        """
        return pulumi.get(self, "compute_expense")

    @compute_expense.setter
    def compute_expense(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "compute_expense", value)

    @property
    @pulumi.getter(name="lastUpdateTime")
    def last_update_time(self) -> Optional[pulumi.Input[str]]:
        """
        Last expense sync time.
        """
        return pulumi.get(self, "last_update_time")

    @last_update_time.setter
    def last_update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_update_time", value)

    @property
    @pulumi.getter
    def message(self) -> Optional[pulumi.Input[str]]:
        """
        Expense sync message if any.
        """
        return pulumi.get(self, "message")

    @message.setter
    def message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message", value)

    @property
    @pulumi.getter(name="networkExpense")
    def network_expense(self) -> Optional[pulumi.Input[float]]:
        """
        Network expense of the entity.
        """
        return pulumi.get(self, "network_expense")

    @network_expense.setter
    def network_expense(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "network_expense", value)

    @property
    @pulumi.getter(name="storageExpense")
    def storage_expense(self) -> Optional[pulumi.Input[float]]:
        """
        Storage expense of the entity.
        """
        return pulumi.get(self, "storage_expense")

    @storage_expense.setter
    def storage_expense(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "storage_expense", value)

    @property
    @pulumi.getter(name="totalExpense")
    def total_expense(self) -> Optional[pulumi.Input[float]]:
        """
        Total expense of the entity.
        """
        return pulumi.get(self, "total_expense")

    @total_expense.setter
    def total_expense(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "total_expense", value)

    @property
    @pulumi.getter
    def unit(self) -> Optional[pulumi.Input[str]]:
        """
        Monetary unit.
        """
        return pulumi.get(self, "unit")

    @unit.setter
    def unit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unit", value)


