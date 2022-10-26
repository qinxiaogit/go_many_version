<?php

$tasks =
    [
        ["tag_name" => ["tag1", "tag2", 'tag3'], "job_name" => "job1"],
        ["tag_name" => ["tag3", "tag4", 'tag5'], "job_name" => "job2"],
        ["tag_name" => ["tag2", 'tag3'], "job_name" => "job3"],
        ["tag_name" => ["tag1", "tag2"], "job_name" => "job4"],
    ];

$taskLen = count($tasks);
/**
 * @desc 计算交集
 */

foreach ($tasks as $k => $task) {
    foreach ($tasks as $index => $item) {
        calculate($task, $item);
    }
}
/***
 * @param array $taskOne
 * @param array $taskTwo
 * @return array
 *  tag1 tag2 tag3 tag4 tag5 tag6 tag7 tag8
 * job1
 *
 */

function calculate(array $taskOne, array $taskTwo): array
{
    print_r(array_intersect($taskOne, $taskTwo));
    return [];
}