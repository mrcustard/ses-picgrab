#!/usr/bin/env ruby


list_of_part_numbers=File.open("NewPartsList.txt")

list_of_part_numbers.each do |part|
    puts "Downloading: " + part
end

