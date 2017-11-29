#!/usr/bin/env bats -t
# Copyright (C) 2017 Miquel Sabaté Solà <msabate@suse.com>
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.

load helpers

function setup() {
    __setup_db
    __source_environment
}

@test "exec works" {
    docker_run portusctl exec rake portus:info
    [[ "${lines[1]}" =~ "Evaluated configuration:" ]]
}

@test "changing the local flag" {
    docker_run portusctl exec -l /srv rake portus:info
    [ $status -eq 1 ]
    [[ "${lines[0]}" =~ "Could not locate Gemfile or .bundle/ directory" ]]
}
