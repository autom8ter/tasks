# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: tasks.proto

require 'google/protobuf'

require 'google/protobuf/empty_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "tasks.IDRequest" do
    optional :id, :int64, 1
  end
  add_message "tasks.Task" do
    optional :id, :int64, 1
    optional :title, :string, 2
    optional :category, :string, 3
    optional :priority, :enum, 4, "tasks.Priority"
    optional :content, :string, 5
    optional :due_date, :string, 6
  end
  add_enum "tasks.Priority" do
    value :Low, 0
    value :Medium, 1
    value :High, 2
  end
end

module Tasks
  IDRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("tasks.IDRequest").msgclass
  Task = Google::Protobuf::DescriptorPool.generated_pool.lookup("tasks.Task").msgclass
  Priority = Google::Protobuf::DescriptorPool.generated_pool.lookup("tasks.Priority").enummodule
end
