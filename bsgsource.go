package main

const bsgHead = "<?xml version=\"1.0\" encoding=\"utf-8\"?>\n<!--Besiege machine save file.-->\n<Machine version=\"1\" bsgVersion=\"1.3\" name=\"%NAME%\">\n    <!--The machine's position and rotation.-->\n    <Global>\n        <Position x=\"0\" y=\"1\" z=\"0\" />\n        <Rotation x=\"0\" y=\"0\" z=\"0\" w=\"1\" />\n    </Global>\n    <!--The machine's additional data or modded data.-->\n    <Data>\n        <StringArray key=\"requiredMods\" />\n    </Data>\n    <!--The machine's blocks.-->\n    <Blocks>"

const bsgTail = "    </Blocks>\n</Machine>"

const transform = "    <Transform>\n        <Position x=\"%PX%\" y=\"%PY%\" z=\"%PZ%\" />\n        <Rotation x=\"1\" y=\"0\" z=\"0\" w=\"-1\" />\n        <Scale x=\"%SX%\" y=\"%SY%\" z=\"%SZ%\" />\n    </Transform>"

const baseBlock = "<Block id=\"0\" guid=\"%GUID%\">\n    %TRANSFORM%\n    <Data />\n</Block>"

const logicBlock = "<Block id=\"68\" guid=\"%GUID%\">\n    %TRANSFORM%\n    <Data>\n        <StringArray key=\"bmt-activate-A\">%INPUT0%</StringArray>\n        <StringArray key=\"bmt-activate-B\">%INPUT1%</StringArray>\n        <Integer key=\"bmt-Gate\">%GATE%</Integer>\n        <Boolean key=\"bmt-toggle-mode\">%TOGGLE%</Boolean>\n        <Boolean key=\"bmt-inverted\">%INVERTED%</Boolean>\n        <StringArray key=\"bmt-emulate\">%OUTPUT%</StringArray>\n    </Data>\n</Block>"

const sensorBlock = "<Block id=\"65\" guid=\"%GUID%\">\n    %TRANSFORM%\n    <Data>\n        <StringArray key=\"bmt-activate\">%INPUT0%</StringArray>\n        <StringArray key=\"bmt-emulate\">%OUTPUT%</StringArray>\n        <Boolean key=\"bmt-non-automatic\">%AUTO%</Boolean>\n        <Boolean key=\"bmt-hold-to-activate\">%HOLD%</Boolean>\n        <Boolean key=\"bmt-reverse\">%REVERSE%</Boolean>\n        <Boolean key=\"bmt-ignore-static\">%STATIC%</Boolean>\n        <Single key=\"bmt-distance\">%DISTANCE%</Single>\n        <Single key=\"bmt-radius\">%RADIUS%</Single>\n    </Data>\n</Block>"

const timerBlock = "<Block id=\"66\" guid=\"%GUID%\">\n    %TRANSFORM%\n    <Data>\n        <StringArray key=\"bmt-activate\">%INPUT0%</StringArray>\n        <StringArray key=\"bmt-emulate\">%OUTPUT%</StringArray>\n        <Boolean key=\"bmt-automatic\">%AUTO%</Boolean>\n        <Boolean key=\"bmt-hold-to-activate\">%HOLD%</Boolean>\n        <Boolean key=\"bmt-can-stop\">%STOP%</Boolean>\n        <Boolean key=\"bmt-loop\">%LOOP%</Boolean>\n        <Single key=\"bmt-wait\">%WAIT%</Single>\n        <Single key=\"bmt-emulation-time\">%TIME%</Single>\n    </Data>\n</Block>"

const altimeterBlock = "<Block id=\"67\" guid=\"%GUID%\">\n    %TRANSFORM%\n    <Data>\n        <StringArray key=\"bmt-activate\">%INPUT0%</StringArray>\n        <StringArray key=\"bmt-emulate\">%OUTPUT%</StringArray>\n        <Boolean key=\"bmt-non-automatic\">%AUTO%</Boolean>\n        <Boolean key=\"bmt-hold-to-activate\">%HOLD%</Boolean>\n        <Boolean key=\"bmt-reverse\">%REVERSE%</Boolean>\n        <Single key=\"bmt-height\">%HEIGHT%</Single>\n    </Data>\n</Block>"

const anglometerBlock = "<Block id=\"69\" guid=\"%GUID%\">\n    %TRANSFORM%\n    <Data>\n        <StringArray key=\"bmt-activate\">%INPUT0%</StringArray>\n        <StringArray key=\"bmt-emulate\">%OUTPUT%</StringArray>\n        <Boolean key=\"bmt-non-automatic\">%AUTO%</Boolean>\n        <Boolean key=\"bmt-hold-to-activate\">%HOLD%</Boolean>\n        <Boolean key=\"bmt-reverse\">%REVERSE%</Boolean>\n        <Single key=\"bmt-start-a\">%START%</Single>\n        <Single key=\"bmt-end-a\">%END%</Single>\n    </Data>\n</Block>"

const speedmeterBlock = "<Block id=\"70\" guid=\"%GUID%\">\n    %TRANSFORM%\n    <Data>\n        <StringArray key=\"bmt-activate\">%INPUT0%</StringArray>\n        <StringArray key=\"bmt-emulate\">%OUTPUT%</StringArray>\n        <Boolean key=\"bmt-non-automatic\">%AUTO%</Boolean>\n        <Boolean key=\"bmt-hold-to-activate\">%HOLD%</Boolean>\n        <Boolean key=\"bmt-reverse\">%REVERSE%</Boolean>\n        <Single key=\"bmt-speed-threshold\">%SPEED%</Single>\n    </Data>\n</Block>"

const msgVar = "<String>None</String>\n<String>Message=%VAR%</String>\n<String>Use=True</String>"
